package objects

// The ReportJob object represents the job status and contains status-specific
// properties for a given report job.
//
// A job with status='error' will populate the message field with an explanation
// of why the job failed. A job with status='complete' will populate the files
// field with a list of file URLs corresponding to the report bundles for the
// specified job id.
//
// The ReportJob object is read-only.
type ReportJob struct {
	// The job's unique identifier.
	JobID *string `json:"jobId,omitempty"`

	// The current status of the job. Values: queued, processing, completed,
	// cancelled, error.
	Status *string `json:"status,omitempty"`

	// The message indicating why the job failed (if status=error).
	Message *string `json:"message,omitempty"`

	// The list of uploaded file URLs for the corresponding job (if status=completed).
	Files []string `json:"files,omitempty"`
}
