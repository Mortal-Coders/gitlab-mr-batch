package model

type GitlabMergeRequest struct {
	Title                 string   `json:"title"`
	SourceBranch          string   `json:"source_branch"`
	TargetBranch          string   `json:"target_branch"`
	TargetProjectId       int      `json:"target_project_id"`
	AssigneeId            int      `json:"assignee_id"`
	AssigneeIds           []int    `json:"assignee_ids"`
	ReviewerIds           []int    `json:"reviewer_ids"`
	Description           string   `json:"description"`
	Labels                []string `json:"labels"`
	AddLabels             []string `json:"add_labels"`
	RemoveLabels          []string `json:"remove_labels"`
	MilestoneId           int      `json:"milestone_id"`
	RemoveSourceBranch    bool     `json:"remove_source_branch"`
	AllowCollaboration    bool     `json:"allow_collaboration"`
	AllowMaintainerToPush bool     `json:"allow_maintainer_to_push"`
	Squash                bool     `json:"squash"`
	MergeAfter            string   `json:"merge_after"`
	ApprovalsBeforeMerge  int      `json:"approvals_before_merge"`
}
