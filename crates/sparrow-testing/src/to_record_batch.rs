use arrow_array::RecordBatch;
use arrow_schema::SchemaRef;
use error_stack::{IntoReport, ResultExt};
use itertools::Itertools;

/// Trait for supporting arbitrary ways of defining a RecordBatch.
pub trait ToRecordBatch {
    /// Convert this to a record batch.
    ///
    /// If provided, the conversion expects to produce the given schema.
    /// This allows controlling how CSV is parsed, etc.
    fn to_record_batch(&self, schema: Option<SchemaRef>) -> super::Result<RecordBatch>;
}

/// Conversion from a JSON-array string to a record batch.
pub struct JsonString<'a>(pub &'a str);

/// Conversion from an array of serde to a record batch (via JSON).
pub struct JsonValues<'a, T: serde::Serialize>(pub &'a [T]);

/// Conversion from a CSV string to a record batch.
pub struct CsvString<'a>(pub &'a str);

impl<'a> ToRecordBatch for JsonString<'a> {
    fn to_record_batch(&self, schema: Option<SchemaRef>) -> super::Result<RecordBatch> {
        // Trim trailing/leading whitespace on each line.
        let json = self.0.lines().map(|line| line.trim()).join("\n");

        // Create the reader
        let reader = std::io::Cursor::new(json.as_bytes());
        let mut reader_builder = arrow_json::ReaderBuilder::new();
        if let Some(schema) = schema {
            reader_builder = reader_builder.with_schema(schema);
        }
        let reader = reader_builder
            .build(reader)
            .into_report()
            .change_context(crate::Error)?;

        // Read all the batches and concatenate them.
        let schema = reader.schema();
        let batches: Vec<_> = reader
            .try_collect()
            .into_report()
            .change_context(crate::Error)?;
        let batch = arrow_select::concat::concat_batches(&schema, &batches)
            .into_report()
            .change_context(crate::Error)?;

        Ok(batch)
    }
}

impl<'a, T: serde::Serialize + std::fmt::Debug + 'a> ToRecordBatch for JsonValues<'a, T> {
    fn to_record_batch(&self, schema: Option<SchemaRef>) -> super::Result<RecordBatch> {
        let value = serde_json::to_value(self.0)
            .into_report()
            .change_context(crate::Error)?;
        let json = match value {
            serde_json::Value::Array(values) => {
                format!("{}", values.iter().format("\n"))
            }
            _ => {
                unreachable!("An array should serialize to an array")
            }
        };

        JsonString(&json).to_record_batch(schema)
    }
}

impl<'a> ToRecordBatch for CsvString<'a> {
    fn to_record_batch(&self, schema: Option<SchemaRef>) -> crate::Result<RecordBatch> {
        // Determine the schema (if not provided).
        let schema = if let Some(schema) = schema {
            schema
        } else {
            let reader = std::io::Cursor::new(self.0.as_bytes());
            let reader = arrow_csv::ReaderBuilder::new()
                .has_header(true)
                .build(reader)
                .into_report()
                .change_context(crate::Error)?;
            reader.schema()
        };

        // Create the reader.
        let reader = std::io::Cursor::new(self.0.as_bytes());
        let reader = arrow_csv::ReaderBuilder::new()
            .has_header(true)
            .with_schema(schema.clone())
            .build(reader)
            .into_report()
            .change_context(crate::Error)?;

        // Read all the batches and concatenate them.
        let batches: Vec<_> = reader
            .try_collect()
            .into_report()
            .change_context(crate::Error)?;
        let batch = arrow_select::concat::concat_batches(&schema, &batches)
            .into_report()
            .change_context(crate::Error)?;

        Ok(batch)
    }
}
