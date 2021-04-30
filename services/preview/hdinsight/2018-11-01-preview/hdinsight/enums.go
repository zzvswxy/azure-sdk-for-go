package hdinsight

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ApplicationState enumerates the values for application state.
type ApplicationState string

const (
	// ACCEPTED ...
	ACCEPTED ApplicationState = "ACCEPTED"
	// FAILED ...
	FAILED ApplicationState = "FAILED"
	// FINISHED ...
	FINISHED ApplicationState = "FINISHED"
	// FINISHING ...
	FINISHING ApplicationState = "FINISHING"
	// KILLED ...
	KILLED ApplicationState = "KILLED"
	// NEW ...
	NEW ApplicationState = "NEW"
	// NEWSAVING ...
	NEWSAVING ApplicationState = "NEW_SAVING"
	// RUNNING ...
	RUNNING ApplicationState = "RUNNING"
	// SUBMITTED ...
	SUBMITTED ApplicationState = "SUBMITTED"
)

// PossibleApplicationStateValues returns an array of possible values for the ApplicationState const type.
func PossibleApplicationStateValues() []ApplicationState {
	return []ApplicationState{ACCEPTED, FAILED, FINISHED, FINISHING, KILLED, NEW, NEWSAVING, RUNNING, SUBMITTED}
}

// JobState enumerates the values for job state.
type JobState string

const (
	// Busy ...
	Busy JobState = "busy"
	// Dead ...
	Dead JobState = "dead"
	// Error ...
	Error JobState = "error"
	// Idle ...
	Idle JobState = "idle"
	// Killed ...
	Killed JobState = "killed"
	// NotStarted ...
	NotStarted JobState = "not_started"
	// Recovering ...
	Recovering JobState = "recovering"
	// Running ...
	Running JobState = "running"
	// ShuttingDown ...
	ShuttingDown JobState = "shutting_down"
	// Starting ...
	Starting JobState = "starting"
	// Success ...
	Success JobState = "success"
)

// PossibleJobStateValues returns an array of possible values for the JobState const type.
func PossibleJobStateValues() []JobState {
	return []JobState{Busy, Dead, Error, Idle, Killed, NotStarted, Recovering, Running, ShuttingDown, Starting, Success}
}

// SessionJobKind enumerates the values for session job kind.
type SessionJobKind string

const (
	// Pyspark ...
	Pyspark SessionJobKind = "pyspark"
	// Spark ...
	Spark SessionJobKind = "spark"
	// Sparkr ...
	Sparkr SessionJobKind = "sparkr"
	// SQL ...
	SQL SessionJobKind = "sql"
)

// PossibleSessionJobKindValues returns an array of possible values for the SessionJobKind const type.
func PossibleSessionJobKindValues() []SessionJobKind {
	return []SessionJobKind{Pyspark, Spark, Sparkr, SQL}
}

// StatementExecutionStatus enumerates the values for statement execution status.
type StatementExecutionStatus string

const (
	// StatementExecutionStatusAbort ...
	StatementExecutionStatusAbort StatementExecutionStatus = "abort"
	// StatementExecutionStatusError ...
	StatementExecutionStatusError StatementExecutionStatus = "error"
	// StatementExecutionStatusOk ...
	StatementExecutionStatusOk StatementExecutionStatus = "ok"
)

// PossibleStatementExecutionStatusValues returns an array of possible values for the StatementExecutionStatus const type.
func PossibleStatementExecutionStatusValues() []StatementExecutionStatus {
	return []StatementExecutionStatus{StatementExecutionStatusAbort, StatementExecutionStatusError, StatementExecutionStatusOk}
}

// StatementState enumerates the values for statement state.
type StatementState string

const (
	// StatementStateAvailable ...
	StatementStateAvailable StatementState = "available"
	// StatementStateCancelled ...
	StatementStateCancelled StatementState = "cancelled"
	// StatementStateCancelling ...
	StatementStateCancelling StatementState = "cancelling"
	// StatementStateError ...
	StatementStateError StatementState = "error"
	// StatementStateRunning ...
	StatementStateRunning StatementState = "running"
	// StatementStateWaiting ...
	StatementStateWaiting StatementState = "waiting"
)

// PossibleStatementStateValues returns an array of possible values for the StatementState const type.
func PossibleStatementStateValues() []StatementState {
	return []StatementState{StatementStateAvailable, StatementStateCancelled, StatementStateCancelling, StatementStateError, StatementStateRunning, StatementStateWaiting}
}