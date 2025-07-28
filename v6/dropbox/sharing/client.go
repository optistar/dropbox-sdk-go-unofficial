// Copyright (c) Dropbox, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package sharing

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"

	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/async"
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/auth"
)

// Client interface describes all routes in this namespace
type Client interface {
	// AddFileMember : Adds specified members to a file.
	AddFileMember(arg *AddFileMemberArgs) (res []*FileMemberActionResult, err error)
	AddFileMemberContext(ctx context.Context, arg *AddFileMemberArgs) (res []*FileMemberActionResult, err error)
	// AddFolderMember : Allows an owner or editor (if the ACL update policy
	// allows) of a shared folder to add another member. For the new member to
	// get access to all the functionality for this folder, you will need to
	// call `mountFolder` on their behalf.
	AddFolderMember(arg *AddFolderMemberArg) (err error)
	AddFolderMemberContext(ctx context.Context, arg *AddFolderMemberArg) (err error)
	// CheckJobStatus : Returns the status of an asynchronous job.
	CheckJobStatus(arg *async.PollArg) (res *JobStatus, err error)
	CheckJobStatusContext(ctx context.Context, arg *async.PollArg) (res *JobStatus, err error)
	// CheckRemoveMemberJobStatus : Returns the status of an asynchronous job
	// for sharing a folder.
	CheckRemoveMemberJobStatus(arg *async.PollArg) (res *RemoveMemberJobStatus, err error)
	CheckRemoveMemberJobStatusContext(ctx context.Context, arg *async.PollArg) (res *RemoveMemberJobStatus, err error)
	// CheckShareJobStatus : Returns the status of an asynchronous job for
	// sharing a folder.
	CheckShareJobStatus(arg *async.PollArg) (res *ShareFolderJobStatus, err error)
	CheckShareJobStatusContext(ctx context.Context, arg *async.PollArg) (res *ShareFolderJobStatus, err error)
	// CreateSharedLink : Create a shared link. If a shared link already exists
	// for the given path, that link is returned. Previously, it was technically
	// possible to break a shared link by moving or renaming the corresponding
	// file or folder. In the future, this will no longer be the case, so your
	// app shouldn't rely on this behavior. Instead, if your app needs to revoke
	// a shared link, use `revokeSharedLink`.
	// Deprecated: Use `CreateSharedLinkWithSettings` instead
	CreateSharedLink(arg *CreateSharedLinkArg) (res *PathLinkMetadata, err error)
	CreateSharedLinkContext(ctx context.Context, arg *CreateSharedLinkArg) (res *PathLinkMetadata, err error)
	// CreateSharedLinkWithSettings : Create a shared link with custom settings.
	// If no settings are given then the default visibility is
	// `RequestedVisibility.public` (The resolved visibility, though, may depend
	// on other aspects such as team and shared folder settings).
	CreateSharedLinkWithSettings(arg *CreateSharedLinkWithSettingsArg) (res IsSharedLinkMetadata, err error)
	CreateSharedLinkWithSettingsContext(ctx context.Context, arg *CreateSharedLinkWithSettingsArg) (res IsSharedLinkMetadata, err error)
	// GetFileMetadata : Returns shared file metadata.
	GetFileMetadata(arg *GetFileMetadataArg) (res *SharedFileMetadata, err error)
	GetFileMetadataContext(ctx context.Context, arg *GetFileMetadataArg) (res *SharedFileMetadata, err error)
	// GetFileMetadataBatch : Returns shared file metadata.
	GetFileMetadataBatch(arg *GetFileMetadataBatchArg) (res []*GetFileMetadataBatchResult, err error)
	GetFileMetadataBatchContext(ctx context.Context, arg *GetFileMetadataBatchArg) (res []*GetFileMetadataBatchResult, err error)
	// GetFolderMetadata : Returns shared folder metadata by its folder ID.
	GetFolderMetadata(arg *GetMetadataArgs) (res *SharedFolderMetadata, err error)
	GetFolderMetadataContext(ctx context.Context, arg *GetMetadataArgs) (res *SharedFolderMetadata, err error)
	// GetSharedLinkFile : Download the shared link's file from a user's
	// Dropbox.
	GetSharedLinkFile(arg *GetSharedLinkMetadataArg) (res IsSharedLinkMetadata, content io.ReadCloser, err error)
	GetSharedLinkFileContext(ctx context.Context, arg *GetSharedLinkMetadataArg) (res IsSharedLinkMetadata, content io.ReadCloser, err error)
	// GetSharedLinkMetadata : Get the shared link's metadata.
	GetSharedLinkMetadata(arg *GetSharedLinkMetadataArg) (res IsSharedLinkMetadata, err error)
	GetSharedLinkMetadataContext(ctx context.Context, arg *GetSharedLinkMetadataArg) (res IsSharedLinkMetadata, err error)
	// GetSharedLinks : Returns a list of `LinkMetadata` objects for this user,
	// including collection links. If no path is given, returns a list of all
	// shared links for the current user, including collection links, up to a
	// maximum of 1000 links. If a non-empty path is given, returns a list of
	// all shared links that allow access to the given path.  Collection links
	// are never returned in this case.
	// Deprecated: Use `ListSharedLinks` instead
	GetSharedLinks(arg *GetSharedLinksArg) (res *GetSharedLinksResult, err error)
	GetSharedLinksContext(ctx context.Context, arg *GetSharedLinksArg) (res *GetSharedLinksResult, err error)
	// ListFileMembers : Use to obtain the members who have been invited to a
	// file, both inherited and uninherited members.
	ListFileMembers(arg *ListFileMembersArg) (res *SharedFileMembers, err error)
	ListFileMembersContext(ctx context.Context, arg *ListFileMembersArg) (res *SharedFileMembers, err error)
	// ListFileMembersBatch : Get members of multiple files at once. The
	// arguments to this route are more limited, and the limit on query result
	// size per file is more strict. To customize the results more, use the
	// individual file endpoint. Inherited users and groups are not included in
	// the result, and permissions are not returned for this endpoint.
	ListFileMembersBatch(arg *ListFileMembersBatchArg) (res []*ListFileMembersBatchResult, err error)
	ListFileMembersBatchContext(ctx context.Context, arg *ListFileMembersBatchArg) (res []*ListFileMembersBatchResult, err error)
	// ListFileMembersContinue : Once a cursor has been retrieved from
	// `listFileMembers` or `listFileMembersBatch`, use this to paginate through
	// all shared file members.
	ListFileMembersContinue(arg *ListFileMembersContinueArg) (res *SharedFileMembers, err error)
	ListFileMembersContinueContext(ctx context.Context, arg *ListFileMembersContinueArg) (res *SharedFileMembers, err error)
	// ListFolderMembers : Returns shared folder membership by its folder ID.
	ListFolderMembers(arg *ListFolderMembersArgs) (res *SharedFolderMembers, err error)
	ListFolderMembersContext(ctx context.Context, arg *ListFolderMembersArgs) (res *SharedFolderMembers, err error)
	// ListFolderMembersContinue : Once a cursor has been retrieved from
	// `listFolderMembers`, use this to paginate through all shared folder
	// members.
	ListFolderMembersContinue(arg *ListFolderMembersContinueArg) (res *SharedFolderMembers, err error)
	ListFolderMembersContinueContext(ctx context.Context, arg *ListFolderMembersContinueArg) (res *SharedFolderMembers, err error)
	// ListFolders : Return the list of all shared folders the current user has
	// access to.
	ListFolders(arg *ListFoldersArgs) (res *ListFoldersResult, err error)
	ListFoldersContext(ctx context.Context, arg *ListFoldersArgs) (res *ListFoldersResult, err error)
	// ListFoldersContinue : Once a cursor has been retrieved from
	// `listFolders`, use this to paginate through all shared folders. The
	// cursor must come from a previous call to `listFolders` or
	// `listFoldersContinue`.
	ListFoldersContinue(arg *ListFoldersContinueArg) (res *ListFoldersResult, err error)
	ListFoldersContinueContext(ctx context.Context, arg *ListFoldersContinueArg) (res *ListFoldersResult, err error)
	// ListMountableFolders : Return the list of all shared folders the current
	// user can mount or unmount.
	ListMountableFolders(arg *ListFoldersArgs) (res *ListFoldersResult, err error)
	ListMountableFoldersContext(ctx context.Context, arg *ListFoldersArgs) (res *ListFoldersResult, err error)
	// ListMountableFoldersContinue : Once a cursor has been retrieved from
	// `listMountableFolders`, use this to paginate through all mountable shared
	// folders. The cursor must come from a previous call to
	// `listMountableFolders` or `listMountableFoldersContinue`.
	ListMountableFoldersContinue(arg *ListFoldersContinueArg) (res *ListFoldersResult, err error)
	ListMountableFoldersContinueContext(ctx context.Context, arg *ListFoldersContinueArg) (res *ListFoldersResult, err error)
	// ListReceivedFiles : Returns a list of all files shared with current user.
	// Does not include files the user has received via shared folders, and does
	// not include unclaimed invitations.
	ListReceivedFiles(arg *ListFilesArg) (res *ListFilesResult, err error)
	ListReceivedFilesContext(ctx context.Context, arg *ListFilesArg) (res *ListFilesResult, err error)
	// ListReceivedFilesContinue : Get more results with a cursor from
	// `listReceivedFiles`.
	ListReceivedFilesContinue(arg *ListFilesContinueArg) (res *ListFilesResult, err error)
	ListReceivedFilesContinueContext(ctx context.Context, arg *ListFilesContinueArg) (res *ListFilesResult, err error)
	// ListSharedLinks : List shared links of this user. If no path is given,
	// returns a list of all shared links for the current user. For members of
	// business teams using team space and member folders, returns all shared
	// links in the team member's home folder unless the team space ID is
	// specified in the request header. For more information, refer to the
	// `Namespace Guide`
	// <https://www.dropbox.com/developers/reference/namespace-guide>. If a
	// non-empty path is given, returns a list of all shared links that allow
	// access to the given path - direct links to the given path and links to
	// parent folders of the given path. Links to parent folders can be
	// suppressed by setting direct_only to true.
	ListSharedLinks(arg *ListSharedLinksArg) (res *ListSharedLinksResult, err error)
	ListSharedLinksContext(ctx context.Context, arg *ListSharedLinksArg) (res *ListSharedLinksResult, err error)
	// ModifySharedLinkSettings : Modify the shared link's settings. If the
	// requested visibility conflict with the shared links policy of the team or
	// the shared folder (in case the linked file is part of a shared folder)
	// then the `LinkPermissions.resolved_visibility` of the returned
	// `SharedLinkMetadata` will reflect the actual visibility of the shared
	// link and the `LinkPermissions.requested_visibility` will reflect the
	// requested visibility.
	ModifySharedLinkSettings(arg *ModifySharedLinkSettingsArgs) (res IsSharedLinkMetadata, err error)
	ModifySharedLinkSettingsContext(ctx context.Context, arg *ModifySharedLinkSettingsArgs) (res IsSharedLinkMetadata, err error)
	// MountFolder : The current user mounts the designated folder. Mount a
	// shared folder for a user after they have been added as a member. Once
	// mounted, the shared folder will appear in their Dropbox.
	MountFolder(arg *MountFolderArg) (res *SharedFolderMetadata, err error)
	MountFolderContext(ctx context.Context, arg *MountFolderArg) (res *SharedFolderMetadata, err error)
	// RelinquishFileMembership : The current user relinquishes their membership
	// in the designated file. Note that the current user may still have
	// inherited access to this file through the parent folder.
	RelinquishFileMembership(arg *RelinquishFileMembershipArg) (err error)
	RelinquishFileMembershipContext(ctx context.Context, arg *RelinquishFileMembershipArg) (err error)
	// RelinquishFolderMembership : The current user relinquishes their
	// membership in the designated shared folder and will no longer have access
	// to the folder.  A folder owner cannot relinquish membership in their own
	// folder. This will run synchronously if leave_a_copy is false, and
	// asynchronously if leave_a_copy is true.
	RelinquishFolderMembership(arg *RelinquishFolderMembershipArg) (res *async.LaunchEmptyResult, err error)
	RelinquishFolderMembershipContext(ctx context.Context, arg *RelinquishFolderMembershipArg) (res *async.LaunchEmptyResult, err error)
	// RemoveFileMember : Identical to remove_file_member_2 but with less
	// information returned.
	// Deprecated: Use `RemoveFileMember2` instead
	RemoveFileMember(arg *RemoveFileMemberArg) (res *FileMemberActionIndividualResult, err error)
	RemoveFileMemberContext(ctx context.Context, arg *RemoveFileMemberArg) (res *FileMemberActionIndividualResult, err error)
	// RemoveFileMember2 : Removes a specified member from the file.
	RemoveFileMember2(arg *RemoveFileMemberArg) (res *FileMemberRemoveActionResult, err error)
	RemoveFileMember2Context(ctx context.Context, arg *RemoveFileMemberArg) (res *FileMemberRemoveActionResult, err error)
	// RemoveFolderMember : Allows an owner or editor (if the ACL update policy
	// allows) of a shared folder to remove another member.
	RemoveFolderMember(arg *RemoveFolderMemberArg) (res *async.LaunchResultBase, err error)
	RemoveFolderMemberContext(ctx context.Context, arg *RemoveFolderMemberArg) (res *async.LaunchResultBase, err error)
	// RevokeSharedLink : Revoke a shared link. Note that even after revoking a
	// shared link to a file, the file may be accessible if there are shared
	// links leading to any of the file parent folders. To list all shared links
	// that enable access to a specific file, you can use the `listSharedLinks`
	// with the file as the `ListSharedLinksArg.path` argument.
	RevokeSharedLink(arg *RevokeSharedLinkArg) (err error)
	RevokeSharedLinkContext(ctx context.Context, arg *RevokeSharedLinkArg) (err error)
	// SetAccessInheritance : Change the inheritance policy of an existing
	// Shared Folder. Only permitted for shared folders in a shared team root.
	// If a `ShareFolderLaunch.async_job_id` is returned, you'll need to call
	// `checkShareJobStatus` until the action completes to get the metadata for
	// the folder.
	SetAccessInheritance(arg *SetAccessInheritanceArg) (res *ShareFolderLaunch, err error)
	SetAccessInheritanceContext(ctx context.Context, arg *SetAccessInheritanceArg) (res *ShareFolderLaunch, err error)
	// ShareFolder : Share a folder with collaborators. Most sharing will be
	// completed synchronously. Large folders will be completed asynchronously.
	// To make testing the async case repeatable, set
	// `ShareFolderArg.force_async`. If a `ShareFolderLaunch.async_job_id` is
	// returned, you'll need to call `checkShareJobStatus` until the action
	// completes to get the metadata for the folder.
	ShareFolder(arg *ShareFolderArg) (res *ShareFolderLaunch, err error)
	ShareFolderContext(ctx context.Context, arg *ShareFolderArg) (res *ShareFolderLaunch, err error)
	// TransferFolder : Transfer ownership of a shared folder to a member of the
	// shared folder. User must have `AccessLevel.owner` access to the shared
	// folder to perform a transfer.
	TransferFolder(arg *TransferFolderArg) (err error)
	TransferFolderContext(ctx context.Context, arg *TransferFolderArg) (err error)
	// UnmountFolder : The current user unmounts the designated folder. They can
	// re-mount the folder at a later time using `mountFolder`.
	UnmountFolder(arg *UnmountFolderArg) (err error)
	UnmountFolderContext(ctx context.Context, arg *UnmountFolderArg) (err error)
	// UnshareFile : Remove all members from this file. Does not remove
	// inherited members.
	UnshareFile(arg *UnshareFileArg) (err error)
	UnshareFileContext(ctx context.Context, arg *UnshareFileArg) (err error)
	// UnshareFolder : Allows a shared folder owner to unshare the folder.
	// You'll need to call `checkJobStatus` to determine if the action has
	// completed successfully.
	UnshareFolder(arg *UnshareFolderArg) (res *async.LaunchEmptyResult, err error)
	UnshareFolderContext(ctx context.Context, arg *UnshareFolderArg) (res *async.LaunchEmptyResult, err error)
	// UpdateFileMember : Changes a member's access on a shared file.
	UpdateFileMember(arg *UpdateFileMemberArgs) (res *MemberAccessLevelResult, err error)
	UpdateFileMemberContext(ctx context.Context, arg *UpdateFileMemberArgs) (res *MemberAccessLevelResult, err error)
	// UpdateFolderMember : Allows an owner or editor of a shared folder to
	// update another member's permissions.
	UpdateFolderMember(arg *UpdateFolderMemberArg) (res *MemberAccessLevelResult, err error)
	UpdateFolderMemberContext(ctx context.Context, arg *UpdateFolderMemberArg) (res *MemberAccessLevelResult, err error)
	// UpdateFolderPolicy : Update the sharing policies for a shared folder.
	// User must have `AccessLevel.owner` access to the shared folder to update
	// its policies.
	UpdateFolderPolicy(arg *UpdateFolderPolicyArg) (res *SharedFolderMetadata, err error)
	UpdateFolderPolicyContext(ctx context.Context, arg *UpdateFolderPolicyArg) (res *SharedFolderMetadata, err error)
}

type apiImpl dropbox.Context

// AddFileMemberAPIError is an error-wrapper for the add_file_member route
type AddFileMemberAPIError struct {
	dropbox.APIError
	EndpointError *AddFileMemberError `json:"error"`
}

func (dbx *apiImpl) AddFileMemberContext(ctx context.Context, arg *AddFileMemberArgs) (res []*FileMemberActionResult, err error) {
	req := dropbox.Request{
		Host:         "api",
		Namespace:    "sharing",
		Route:        "add_file_member",
		Auth:         "user",
		Style:        "rpc",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr AddFileMemberAPIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return
	}

	_ = respBody
	return
}

func (dbx *apiImpl) AddFileMember(arg *AddFileMemberArgs) (res []*FileMemberActionResult, err error) {
	return dbx.AddFileMemberContext(context.Background(), arg)
}

// AddFolderMemberAPIError is an error-wrapper for the add_folder_member route
type AddFolderMemberAPIError struct {
	dropbox.APIError
	EndpointError *AddFolderMemberError `json:"error"`
}

func (dbx *apiImpl) AddFolderMemberContext(ctx context.Context, arg *AddFolderMemberArg) (err error) {
	req := dropbox.Request{
		Host:         "api",
		Namespace:    "sharing",
		Route:        "add_folder_member",
		Auth:         "user",
		Style:        "rpc",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr AddFolderMemberAPIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	_ = resp
	_ = respBody
	return
}

func (dbx *apiImpl) AddFolderMember(arg *AddFolderMemberArg) (err error) {
	return dbx.AddFolderMemberContext(context.Background(), arg)
}

// CheckJobStatusAPIError is an error-wrapper for the check_job_status route
type CheckJobStatusAPIError struct {
	dropbox.APIError
	EndpointError *async.PollError `json:"error"`
}

func (dbx *apiImpl) CheckJobStatusContext(ctx context.Context, arg *async.PollArg) (res *JobStatus, err error) {
	req := dropbox.Request{
		Host:         "api",
		Namespace:    "sharing",
		Route:        "check_job_status",
		Auth:         "user",
		Style:        "rpc",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr CheckJobStatusAPIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return
	}

	_ = respBody
	return
}

func (dbx *apiImpl) CheckJobStatus(arg *async.PollArg) (res *JobStatus, err error) {
	return dbx.CheckJobStatusContext(context.Background(), arg)
}

// CheckRemoveMemberJobStatusAPIError is an error-wrapper for the check_remove_member_job_status route
type CheckRemoveMemberJobStatusAPIError struct {
	dropbox.APIError
	EndpointError *async.PollError `json:"error"`
}

func (dbx *apiImpl) CheckRemoveMemberJobStatusContext(ctx context.Context, arg *async.PollArg) (res *RemoveMemberJobStatus, err error) {
	req := dropbox.Request{
		Host:         "api",
		Namespace:    "sharing",
		Route:        "check_remove_member_job_status",
		Auth:         "user",
		Style:        "rpc",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr CheckRemoveMemberJobStatusAPIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return
	}

	_ = respBody
	return
}

func (dbx *apiImpl) CheckRemoveMemberJobStatus(arg *async.PollArg) (res *RemoveMemberJobStatus, err error) {
	return dbx.CheckRemoveMemberJobStatusContext(context.Background(), arg)
}

// CheckShareJobStatusAPIError is an error-wrapper for the check_share_job_status route
type CheckShareJobStatusAPIError struct {
	dropbox.APIError
	EndpointError *async.PollError `json:"error"`
}

func (dbx *apiImpl) CheckShareJobStatusContext(ctx context.Context, arg *async.PollArg) (res *ShareFolderJobStatus, err error) {
	req := dropbox.Request{
		Host:         "api",
		Namespace:    "sharing",
		Route:        "check_share_job_status",
		Auth:         "user",
		Style:        "rpc",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr CheckShareJobStatusAPIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return
	}

	_ = respBody
	return
}

func (dbx *apiImpl) CheckShareJobStatus(arg *async.PollArg) (res *ShareFolderJobStatus, err error) {
	return dbx.CheckShareJobStatusContext(context.Background(), arg)
}

// CreateSharedLinkAPIError is an error-wrapper for the create_shared_link route
type CreateSharedLinkAPIError struct {
	dropbox.APIError
	EndpointError *CreateSharedLinkError `json:"error"`
}

func (dbx *apiImpl) CreateSharedLinkContext(ctx context.Context, arg *CreateSharedLinkArg) (res *PathLinkMetadata, err error) {
	log.Printf("WARNING: API `CreateSharedLink` is deprecated")
	log.Printf("Use API `CreateSharedLinkWithSettings` instead")

	req := dropbox.Request{
		Host:         "api",
		Namespace:    "sharing",
		Route:        "create_shared_link",
		Auth:         "user",
		Style:        "rpc",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr CreateSharedLinkAPIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return
	}

	_ = respBody
	return
}

func (dbx *apiImpl) CreateSharedLink(arg *CreateSharedLinkArg) (res *PathLinkMetadata, err error) {
	return dbx.CreateSharedLinkContext(context.Background(), arg)
}

// CreateSharedLinkWithSettingsAPIError is an error-wrapper for the create_shared_link_with_settings route
type CreateSharedLinkWithSettingsAPIError struct {
	dropbox.APIError
	EndpointError *CreateSharedLinkWithSettingsError `json:"error"`
}

func (dbx *apiImpl) CreateSharedLinkWithSettingsContext(ctx context.Context, arg *CreateSharedLinkWithSettingsArg) (res IsSharedLinkMetadata, err error) {
	req := dropbox.Request{
		Host:         "api",
		Namespace:    "sharing",
		Route:        "create_shared_link_with_settings",
		Auth:         "user",
		Style:        "rpc",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr CreateSharedLinkWithSettingsAPIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	var tmp sharedLinkMetadataUnion
	err = json.Unmarshal(resp, &tmp)
	if err != nil {
		return
	}
	switch tmp.Tag {
	case "file":
		res = tmp.File

	case "folder":
		res = tmp.Folder

	}
	_ = respBody
	return
}

func (dbx *apiImpl) CreateSharedLinkWithSettings(arg *CreateSharedLinkWithSettingsArg) (res IsSharedLinkMetadata, err error) {
	return dbx.CreateSharedLinkWithSettingsContext(context.Background(), arg)
}

// GetFileMetadataAPIError is an error-wrapper for the get_file_metadata route
type GetFileMetadataAPIError struct {
	dropbox.APIError
	EndpointError *GetFileMetadataError `json:"error"`
}

func (dbx *apiImpl) GetFileMetadataContext(ctx context.Context, arg *GetFileMetadataArg) (res *SharedFileMetadata, err error) {
	req := dropbox.Request{
		Host:         "api",
		Namespace:    "sharing",
		Route:        "get_file_metadata",
		Auth:         "user",
		Style:        "rpc",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr GetFileMetadataAPIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return
	}

	_ = respBody
	return
}

func (dbx *apiImpl) GetFileMetadata(arg *GetFileMetadataArg) (res *SharedFileMetadata, err error) {
	return dbx.GetFileMetadataContext(context.Background(), arg)
}

// GetFileMetadataBatchAPIError is an error-wrapper for the get_file_metadata/batch route
type GetFileMetadataBatchAPIError struct {
	dropbox.APIError
	EndpointError *SharingUserError `json:"error"`
}

func (dbx *apiImpl) GetFileMetadataBatchContext(ctx context.Context, arg *GetFileMetadataBatchArg) (res []*GetFileMetadataBatchResult, err error) {
	req := dropbox.Request{
		Host:         "api",
		Namespace:    "sharing",
		Route:        "get_file_metadata/batch",
		Auth:         "user",
		Style:        "rpc",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr GetFileMetadataBatchAPIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return
	}

	_ = respBody
	return
}

func (dbx *apiImpl) GetFileMetadataBatch(arg *GetFileMetadataBatchArg) (res []*GetFileMetadataBatchResult, err error) {
	return dbx.GetFileMetadataBatchContext(context.Background(), arg)
}

// GetFolderMetadataAPIError is an error-wrapper for the get_folder_metadata route
type GetFolderMetadataAPIError struct {
	dropbox.APIError
	EndpointError *SharedFolderAccessError `json:"error"`
}

func (dbx *apiImpl) GetFolderMetadataContext(ctx context.Context, arg *GetMetadataArgs) (res *SharedFolderMetadata, err error) {
	req := dropbox.Request{
		Host:         "api",
		Namespace:    "sharing",
		Route:        "get_folder_metadata",
		Auth:         "user",
		Style:        "rpc",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr GetFolderMetadataAPIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return
	}

	_ = respBody
	return
}

func (dbx *apiImpl) GetFolderMetadata(arg *GetMetadataArgs) (res *SharedFolderMetadata, err error) {
	return dbx.GetFolderMetadataContext(context.Background(), arg)
}

// GetSharedLinkFileAPIError is an error-wrapper for the get_shared_link_file route
type GetSharedLinkFileAPIError struct {
	dropbox.APIError
	EndpointError *GetSharedLinkFileError `json:"error"`
}

func (dbx *apiImpl) GetSharedLinkFileContext(ctx context.Context, arg *GetSharedLinkMetadataArg) (res IsSharedLinkMetadata, content io.ReadCloser, err error) {
	req := dropbox.Request{
		Host:         "content",
		Namespace:    "sharing",
		Route:        "get_shared_link_file",
		Auth:         "user",
		Style:        "download",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr GetSharedLinkFileAPIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	var tmp sharedLinkMetadataUnion
	err = json.Unmarshal(resp, &tmp)
	if err != nil {
		return
	}
	switch tmp.Tag {
	case "file":
		res = tmp.File

	case "folder":
		res = tmp.Folder

	}
	content = respBody
	return
}

func (dbx *apiImpl) GetSharedLinkFile(arg *GetSharedLinkMetadataArg) (res IsSharedLinkMetadata, content io.ReadCloser, err error) {
	return dbx.GetSharedLinkFileContext(context.Background(), arg)
}

// GetSharedLinkMetadataAPIError is an error-wrapper for the get_shared_link_metadata route
type GetSharedLinkMetadataAPIError struct {
	dropbox.APIError
	EndpointError *SharedLinkError `json:"error"`
}

func (dbx *apiImpl) GetSharedLinkMetadataContext(ctx context.Context, arg *GetSharedLinkMetadataArg) (res IsSharedLinkMetadata, err error) {
	req := dropbox.Request{
		Host:         "api",
		Namespace:    "sharing",
		Route:        "get_shared_link_metadata",
		Auth:         "app, user",
		Style:        "rpc",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr GetSharedLinkMetadataAPIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	var tmp sharedLinkMetadataUnion
	err = json.Unmarshal(resp, &tmp)
	if err != nil {
		return
	}
	switch tmp.Tag {
	case "file":
		res = tmp.File

	case "folder":
		res = tmp.Folder

	}
	_ = respBody
	return
}

func (dbx *apiImpl) GetSharedLinkMetadata(arg *GetSharedLinkMetadataArg) (res IsSharedLinkMetadata, err error) {
	return dbx.GetSharedLinkMetadataContext(context.Background(), arg)
}

// GetSharedLinksAPIError is an error-wrapper for the get_shared_links route
type GetSharedLinksAPIError struct {
	dropbox.APIError
	EndpointError *GetSharedLinksError `json:"error"`
}

func (dbx *apiImpl) GetSharedLinksContext(ctx context.Context, arg *GetSharedLinksArg) (res *GetSharedLinksResult, err error) {
	log.Printf("WARNING: API `GetSharedLinks` is deprecated")
	log.Printf("Use API `ListSharedLinks` instead")

	req := dropbox.Request{
		Host:         "api",
		Namespace:    "sharing",
		Route:        "get_shared_links",
		Auth:         "user",
		Style:        "rpc",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr GetSharedLinksAPIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return
	}

	_ = respBody
	return
}

func (dbx *apiImpl) GetSharedLinks(arg *GetSharedLinksArg) (res *GetSharedLinksResult, err error) {
	return dbx.GetSharedLinksContext(context.Background(), arg)
}

// ListFileMembersAPIError is an error-wrapper for the list_file_members route
type ListFileMembersAPIError struct {
	dropbox.APIError
	EndpointError *ListFileMembersError `json:"error"`
}

func (dbx *apiImpl) ListFileMembersContext(ctx context.Context, arg *ListFileMembersArg) (res *SharedFileMembers, err error) {
	req := dropbox.Request{
		Host:         "api",
		Namespace:    "sharing",
		Route:        "list_file_members",
		Auth:         "user",
		Style:        "rpc",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr ListFileMembersAPIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return
	}

	_ = respBody
	return
}

func (dbx *apiImpl) ListFileMembers(arg *ListFileMembersArg) (res *SharedFileMembers, err error) {
	return dbx.ListFileMembersContext(context.Background(), arg)
}

// ListFileMembersBatchAPIError is an error-wrapper for the list_file_members/batch route
type ListFileMembersBatchAPIError struct {
	dropbox.APIError
	EndpointError *SharingUserError `json:"error"`
}

func (dbx *apiImpl) ListFileMembersBatchContext(ctx context.Context, arg *ListFileMembersBatchArg) (res []*ListFileMembersBatchResult, err error) {
	req := dropbox.Request{
		Host:         "api",
		Namespace:    "sharing",
		Route:        "list_file_members/batch",
		Auth:         "user",
		Style:        "rpc",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr ListFileMembersBatchAPIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return
	}

	_ = respBody
	return
}

func (dbx *apiImpl) ListFileMembersBatch(arg *ListFileMembersBatchArg) (res []*ListFileMembersBatchResult, err error) {
	return dbx.ListFileMembersBatchContext(context.Background(), arg)
}

// ListFileMembersContinueAPIError is an error-wrapper for the list_file_members/continue route
type ListFileMembersContinueAPIError struct {
	dropbox.APIError
	EndpointError *ListFileMembersContinueError `json:"error"`
}

func (dbx *apiImpl) ListFileMembersContinueContext(ctx context.Context, arg *ListFileMembersContinueArg) (res *SharedFileMembers, err error) {
	req := dropbox.Request{
		Host:         "api",
		Namespace:    "sharing",
		Route:        "list_file_members/continue",
		Auth:         "user",
		Style:        "rpc",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr ListFileMembersContinueAPIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return
	}

	_ = respBody
	return
}

func (dbx *apiImpl) ListFileMembersContinue(arg *ListFileMembersContinueArg) (res *SharedFileMembers, err error) {
	return dbx.ListFileMembersContinueContext(context.Background(), arg)
}

// ListFolderMembersAPIError is an error-wrapper for the list_folder_members route
type ListFolderMembersAPIError struct {
	dropbox.APIError
	EndpointError *SharedFolderAccessError `json:"error"`
}

func (dbx *apiImpl) ListFolderMembersContext(ctx context.Context, arg *ListFolderMembersArgs) (res *SharedFolderMembers, err error) {
	req := dropbox.Request{
		Host:         "api",
		Namespace:    "sharing",
		Route:        "list_folder_members",
		Auth:         "user",
		Style:        "rpc",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr ListFolderMembersAPIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return
	}

	_ = respBody
	return
}

func (dbx *apiImpl) ListFolderMembers(arg *ListFolderMembersArgs) (res *SharedFolderMembers, err error) {
	return dbx.ListFolderMembersContext(context.Background(), arg)
}

// ListFolderMembersContinueAPIError is an error-wrapper for the list_folder_members/continue route
type ListFolderMembersContinueAPIError struct {
	dropbox.APIError
	EndpointError *ListFolderMembersContinueError `json:"error"`
}

func (dbx *apiImpl) ListFolderMembersContinueContext(ctx context.Context, arg *ListFolderMembersContinueArg) (res *SharedFolderMembers, err error) {
	req := dropbox.Request{
		Host:         "api",
		Namespace:    "sharing",
		Route:        "list_folder_members/continue",
		Auth:         "user",
		Style:        "rpc",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr ListFolderMembersContinueAPIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return
	}

	_ = respBody
	return
}

func (dbx *apiImpl) ListFolderMembersContinue(arg *ListFolderMembersContinueArg) (res *SharedFolderMembers, err error) {
	return dbx.ListFolderMembersContinueContext(context.Background(), arg)
}

// ListFoldersAPIError is an error-wrapper for the list_folders route
type ListFoldersAPIError struct {
	dropbox.APIError
	EndpointError struct{} `json:"error"`
}

func (dbx *apiImpl) ListFoldersContext(ctx context.Context, arg *ListFoldersArgs) (res *ListFoldersResult, err error) {
	req := dropbox.Request{
		Host:         "api",
		Namespace:    "sharing",
		Route:        "list_folders",
		Auth:         "user",
		Style:        "rpc",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr ListFoldersAPIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return
	}

	_ = respBody
	return
}

func (dbx *apiImpl) ListFolders(arg *ListFoldersArgs) (res *ListFoldersResult, err error) {
	return dbx.ListFoldersContext(context.Background(), arg)
}

// ListFoldersContinueAPIError is an error-wrapper for the list_folders/continue route
type ListFoldersContinueAPIError struct {
	dropbox.APIError
	EndpointError *ListFoldersContinueError `json:"error"`
}

func (dbx *apiImpl) ListFoldersContinueContext(ctx context.Context, arg *ListFoldersContinueArg) (res *ListFoldersResult, err error) {
	req := dropbox.Request{
		Host:         "api",
		Namespace:    "sharing",
		Route:        "list_folders/continue",
		Auth:         "user",
		Style:        "rpc",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr ListFoldersContinueAPIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return
	}

	_ = respBody
	return
}

func (dbx *apiImpl) ListFoldersContinue(arg *ListFoldersContinueArg) (res *ListFoldersResult, err error) {
	return dbx.ListFoldersContinueContext(context.Background(), arg)
}

// ListMountableFoldersAPIError is an error-wrapper for the list_mountable_folders route
type ListMountableFoldersAPIError struct {
	dropbox.APIError
	EndpointError struct{} `json:"error"`
}

func (dbx *apiImpl) ListMountableFoldersContext(ctx context.Context, arg *ListFoldersArgs) (res *ListFoldersResult, err error) {
	req := dropbox.Request{
		Host:         "api",
		Namespace:    "sharing",
		Route:        "list_mountable_folders",
		Auth:         "user",
		Style:        "rpc",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr ListMountableFoldersAPIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return
	}

	_ = respBody
	return
}

func (dbx *apiImpl) ListMountableFolders(arg *ListFoldersArgs) (res *ListFoldersResult, err error) {
	return dbx.ListMountableFoldersContext(context.Background(), arg)
}

// ListMountableFoldersContinueAPIError is an error-wrapper for the list_mountable_folders/continue route
type ListMountableFoldersContinueAPIError struct {
	dropbox.APIError
	EndpointError *ListFoldersContinueError `json:"error"`
}

func (dbx *apiImpl) ListMountableFoldersContinueContext(ctx context.Context, arg *ListFoldersContinueArg) (res *ListFoldersResult, err error) {
	req := dropbox.Request{
		Host:         "api",
		Namespace:    "sharing",
		Route:        "list_mountable_folders/continue",
		Auth:         "user",
		Style:        "rpc",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr ListMountableFoldersContinueAPIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return
	}

	_ = respBody
	return
}

func (dbx *apiImpl) ListMountableFoldersContinue(arg *ListFoldersContinueArg) (res *ListFoldersResult, err error) {
	return dbx.ListMountableFoldersContinueContext(context.Background(), arg)
}

// ListReceivedFilesAPIError is an error-wrapper for the list_received_files route
type ListReceivedFilesAPIError struct {
	dropbox.APIError
	EndpointError *SharingUserError `json:"error"`
}

func (dbx *apiImpl) ListReceivedFilesContext(ctx context.Context, arg *ListFilesArg) (res *ListFilesResult, err error) {
	req := dropbox.Request{
		Host:         "api",
		Namespace:    "sharing",
		Route:        "list_received_files",
		Auth:         "user",
		Style:        "rpc",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr ListReceivedFilesAPIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return
	}

	_ = respBody
	return
}

func (dbx *apiImpl) ListReceivedFiles(arg *ListFilesArg) (res *ListFilesResult, err error) {
	return dbx.ListReceivedFilesContext(context.Background(), arg)
}

// ListReceivedFilesContinueAPIError is an error-wrapper for the list_received_files/continue route
type ListReceivedFilesContinueAPIError struct {
	dropbox.APIError
	EndpointError *ListFilesContinueError `json:"error"`
}

func (dbx *apiImpl) ListReceivedFilesContinueContext(ctx context.Context, arg *ListFilesContinueArg) (res *ListFilesResult, err error) {
	req := dropbox.Request{
		Host:         "api",
		Namespace:    "sharing",
		Route:        "list_received_files/continue",
		Auth:         "user",
		Style:        "rpc",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr ListReceivedFilesContinueAPIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return
	}

	_ = respBody
	return
}

func (dbx *apiImpl) ListReceivedFilesContinue(arg *ListFilesContinueArg) (res *ListFilesResult, err error) {
	return dbx.ListReceivedFilesContinueContext(context.Background(), arg)
}

// ListSharedLinksAPIError is an error-wrapper for the list_shared_links route
type ListSharedLinksAPIError struct {
	dropbox.APIError
	EndpointError *ListSharedLinksError `json:"error"`
}

func (dbx *apiImpl) ListSharedLinksContext(ctx context.Context, arg *ListSharedLinksArg) (res *ListSharedLinksResult, err error) {
	req := dropbox.Request{
		Host:         "api",
		Namespace:    "sharing",
		Route:        "list_shared_links",
		Auth:         "user",
		Style:        "rpc",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr ListSharedLinksAPIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return
	}

	_ = respBody
	return
}

func (dbx *apiImpl) ListSharedLinks(arg *ListSharedLinksArg) (res *ListSharedLinksResult, err error) {
	return dbx.ListSharedLinksContext(context.Background(), arg)
}

// ModifySharedLinkSettingsAPIError is an error-wrapper for the modify_shared_link_settings route
type ModifySharedLinkSettingsAPIError struct {
	dropbox.APIError
	EndpointError *ModifySharedLinkSettingsError `json:"error"`
}

func (dbx *apiImpl) ModifySharedLinkSettingsContext(ctx context.Context, arg *ModifySharedLinkSettingsArgs) (res IsSharedLinkMetadata, err error) {
	req := dropbox.Request{
		Host:         "api",
		Namespace:    "sharing",
		Route:        "modify_shared_link_settings",
		Auth:         "user",
		Style:        "rpc",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr ModifySharedLinkSettingsAPIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	var tmp sharedLinkMetadataUnion
	err = json.Unmarshal(resp, &tmp)
	if err != nil {
		return
	}
	switch tmp.Tag {
	case "file":
		res = tmp.File

	case "folder":
		res = tmp.Folder

	}
	_ = respBody
	return
}

func (dbx *apiImpl) ModifySharedLinkSettings(arg *ModifySharedLinkSettingsArgs) (res IsSharedLinkMetadata, err error) {
	return dbx.ModifySharedLinkSettingsContext(context.Background(), arg)
}

// MountFolderAPIError is an error-wrapper for the mount_folder route
type MountFolderAPIError struct {
	dropbox.APIError
	EndpointError *MountFolderError `json:"error"`
}

func (dbx *apiImpl) MountFolderContext(ctx context.Context, arg *MountFolderArg) (res *SharedFolderMetadata, err error) {
	req := dropbox.Request{
		Host:         "api",
		Namespace:    "sharing",
		Route:        "mount_folder",
		Auth:         "user",
		Style:        "rpc",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr MountFolderAPIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return
	}

	_ = respBody
	return
}

func (dbx *apiImpl) MountFolder(arg *MountFolderArg) (res *SharedFolderMetadata, err error) {
	return dbx.MountFolderContext(context.Background(), arg)
}

// RelinquishFileMembershipAPIError is an error-wrapper for the relinquish_file_membership route
type RelinquishFileMembershipAPIError struct {
	dropbox.APIError
	EndpointError *RelinquishFileMembershipError `json:"error"`
}

func (dbx *apiImpl) RelinquishFileMembershipContext(ctx context.Context, arg *RelinquishFileMembershipArg) (err error) {
	req := dropbox.Request{
		Host:         "api",
		Namespace:    "sharing",
		Route:        "relinquish_file_membership",
		Auth:         "user",
		Style:        "rpc",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr RelinquishFileMembershipAPIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	_ = resp
	_ = respBody
	return
}

func (dbx *apiImpl) RelinquishFileMembership(arg *RelinquishFileMembershipArg) (err error) {
	return dbx.RelinquishFileMembershipContext(context.Background(), arg)
}

// RelinquishFolderMembershipAPIError is an error-wrapper for the relinquish_folder_membership route
type RelinquishFolderMembershipAPIError struct {
	dropbox.APIError
	EndpointError *RelinquishFolderMembershipError `json:"error"`
}

func (dbx *apiImpl) RelinquishFolderMembershipContext(ctx context.Context, arg *RelinquishFolderMembershipArg) (res *async.LaunchEmptyResult, err error) {
	req := dropbox.Request{
		Host:         "api",
		Namespace:    "sharing",
		Route:        "relinquish_folder_membership",
		Auth:         "user",
		Style:        "rpc",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr RelinquishFolderMembershipAPIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return
	}

	_ = respBody
	return
}

func (dbx *apiImpl) RelinquishFolderMembership(arg *RelinquishFolderMembershipArg) (res *async.LaunchEmptyResult, err error) {
	return dbx.RelinquishFolderMembershipContext(context.Background(), arg)
}

// RemoveFileMemberAPIError is an error-wrapper for the remove_file_member route
type RemoveFileMemberAPIError struct {
	dropbox.APIError
	EndpointError *RemoveFileMemberError `json:"error"`
}

func (dbx *apiImpl) RemoveFileMemberContext(ctx context.Context, arg *RemoveFileMemberArg) (res *FileMemberActionIndividualResult, err error) {
	log.Printf("WARNING: API `RemoveFileMember` is deprecated")
	log.Printf("Use API `RemoveFileMember2` instead")

	req := dropbox.Request{
		Host:         "api",
		Namespace:    "sharing",
		Route:        "remove_file_member",
		Auth:         "user",
		Style:        "rpc",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr RemoveFileMemberAPIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return
	}

	_ = respBody
	return
}

func (dbx *apiImpl) RemoveFileMember(arg *RemoveFileMemberArg) (res *FileMemberActionIndividualResult, err error) {
	return dbx.RemoveFileMemberContext(context.Background(), arg)
}

// RemoveFileMember2APIError is an error-wrapper for the remove_file_member_2 route
type RemoveFileMember2APIError struct {
	dropbox.APIError
	EndpointError *RemoveFileMemberError `json:"error"`
}

func (dbx *apiImpl) RemoveFileMember2Context(ctx context.Context, arg *RemoveFileMemberArg) (res *FileMemberRemoveActionResult, err error) {
	req := dropbox.Request{
		Host:         "api",
		Namespace:    "sharing",
		Route:        "remove_file_member_2",
		Auth:         "user",
		Style:        "rpc",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr RemoveFileMember2APIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return
	}

	_ = respBody
	return
}

func (dbx *apiImpl) RemoveFileMember2(arg *RemoveFileMemberArg) (res *FileMemberRemoveActionResult, err error) {
	return dbx.RemoveFileMember2Context(context.Background(), arg)
}

// RemoveFolderMemberAPIError is an error-wrapper for the remove_folder_member route
type RemoveFolderMemberAPIError struct {
	dropbox.APIError
	EndpointError *RemoveFolderMemberError `json:"error"`
}

func (dbx *apiImpl) RemoveFolderMemberContext(ctx context.Context, arg *RemoveFolderMemberArg) (res *async.LaunchResultBase, err error) {
	req := dropbox.Request{
		Host:         "api",
		Namespace:    "sharing",
		Route:        "remove_folder_member",
		Auth:         "user",
		Style:        "rpc",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr RemoveFolderMemberAPIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return
	}

	_ = respBody
	return
}

func (dbx *apiImpl) RemoveFolderMember(arg *RemoveFolderMemberArg) (res *async.LaunchResultBase, err error) {
	return dbx.RemoveFolderMemberContext(context.Background(), arg)
}

// RevokeSharedLinkAPIError is an error-wrapper for the revoke_shared_link route
type RevokeSharedLinkAPIError struct {
	dropbox.APIError
	EndpointError *RevokeSharedLinkError `json:"error"`
}

func (dbx *apiImpl) RevokeSharedLinkContext(ctx context.Context, arg *RevokeSharedLinkArg) (err error) {
	req := dropbox.Request{
		Host:         "api",
		Namespace:    "sharing",
		Route:        "revoke_shared_link",
		Auth:         "user",
		Style:        "rpc",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr RevokeSharedLinkAPIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	_ = resp
	_ = respBody
	return
}

func (dbx *apiImpl) RevokeSharedLink(arg *RevokeSharedLinkArg) (err error) {
	return dbx.RevokeSharedLinkContext(context.Background(), arg)
}

// SetAccessInheritanceAPIError is an error-wrapper for the set_access_inheritance route
type SetAccessInheritanceAPIError struct {
	dropbox.APIError
	EndpointError *SetAccessInheritanceError `json:"error"`
}

func (dbx *apiImpl) SetAccessInheritanceContext(ctx context.Context, arg *SetAccessInheritanceArg) (res *ShareFolderLaunch, err error) {
	req := dropbox.Request{
		Host:         "api",
		Namespace:    "sharing",
		Route:        "set_access_inheritance",
		Auth:         "user",
		Style:        "rpc",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr SetAccessInheritanceAPIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return
	}

	_ = respBody
	return
}

func (dbx *apiImpl) SetAccessInheritance(arg *SetAccessInheritanceArg) (res *ShareFolderLaunch, err error) {
	return dbx.SetAccessInheritanceContext(context.Background(), arg)
}

// ShareFolderAPIError is an error-wrapper for the share_folder route
type ShareFolderAPIError struct {
	dropbox.APIError
	EndpointError *ShareFolderError `json:"error"`
}

func (dbx *apiImpl) ShareFolderContext(ctx context.Context, arg *ShareFolderArg) (res *ShareFolderLaunch, err error) {
	req := dropbox.Request{
		Host:         "api",
		Namespace:    "sharing",
		Route:        "share_folder",
		Auth:         "user",
		Style:        "rpc",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr ShareFolderAPIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return
	}

	_ = respBody
	return
}

func (dbx *apiImpl) ShareFolder(arg *ShareFolderArg) (res *ShareFolderLaunch, err error) {
	return dbx.ShareFolderContext(context.Background(), arg)
}

// TransferFolderAPIError is an error-wrapper for the transfer_folder route
type TransferFolderAPIError struct {
	dropbox.APIError
	EndpointError *TransferFolderError `json:"error"`
}

func (dbx *apiImpl) TransferFolderContext(ctx context.Context, arg *TransferFolderArg) (err error) {
	req := dropbox.Request{
		Host:         "api",
		Namespace:    "sharing",
		Route:        "transfer_folder",
		Auth:         "user",
		Style:        "rpc",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr TransferFolderAPIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	_ = resp
	_ = respBody
	return
}

func (dbx *apiImpl) TransferFolder(arg *TransferFolderArg) (err error) {
	return dbx.TransferFolderContext(context.Background(), arg)
}

// UnmountFolderAPIError is an error-wrapper for the unmount_folder route
type UnmountFolderAPIError struct {
	dropbox.APIError
	EndpointError *UnmountFolderError `json:"error"`
}

func (dbx *apiImpl) UnmountFolderContext(ctx context.Context, arg *UnmountFolderArg) (err error) {
	req := dropbox.Request{
		Host:         "api",
		Namespace:    "sharing",
		Route:        "unmount_folder",
		Auth:         "user",
		Style:        "rpc",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr UnmountFolderAPIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	_ = resp
	_ = respBody
	return
}

func (dbx *apiImpl) UnmountFolder(arg *UnmountFolderArg) (err error) {
	return dbx.UnmountFolderContext(context.Background(), arg)
}

// UnshareFileAPIError is an error-wrapper for the unshare_file route
type UnshareFileAPIError struct {
	dropbox.APIError
	EndpointError *UnshareFileError `json:"error"`
}

func (dbx *apiImpl) UnshareFileContext(ctx context.Context, arg *UnshareFileArg) (err error) {
	req := dropbox.Request{
		Host:         "api",
		Namespace:    "sharing",
		Route:        "unshare_file",
		Auth:         "user",
		Style:        "rpc",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr UnshareFileAPIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	_ = resp
	_ = respBody
	return
}

func (dbx *apiImpl) UnshareFile(arg *UnshareFileArg) (err error) {
	return dbx.UnshareFileContext(context.Background(), arg)
}

// UnshareFolderAPIError is an error-wrapper for the unshare_folder route
type UnshareFolderAPIError struct {
	dropbox.APIError
	EndpointError *UnshareFolderError `json:"error"`
}

func (dbx *apiImpl) UnshareFolderContext(ctx context.Context, arg *UnshareFolderArg) (res *async.LaunchEmptyResult, err error) {
	req := dropbox.Request{
		Host:         "api",
		Namespace:    "sharing",
		Route:        "unshare_folder",
		Auth:         "user",
		Style:        "rpc",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr UnshareFolderAPIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return
	}

	_ = respBody
	return
}

func (dbx *apiImpl) UnshareFolder(arg *UnshareFolderArg) (res *async.LaunchEmptyResult, err error) {
	return dbx.UnshareFolderContext(context.Background(), arg)
}

// UpdateFileMemberAPIError is an error-wrapper for the update_file_member route
type UpdateFileMemberAPIError struct {
	dropbox.APIError
	EndpointError *FileMemberActionError `json:"error"`
}

func (dbx *apiImpl) UpdateFileMemberContext(ctx context.Context, arg *UpdateFileMemberArgs) (res *MemberAccessLevelResult, err error) {
	req := dropbox.Request{
		Host:         "api",
		Namespace:    "sharing",
		Route:        "update_file_member",
		Auth:         "user",
		Style:        "rpc",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr UpdateFileMemberAPIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return
	}

	_ = respBody
	return
}

func (dbx *apiImpl) UpdateFileMember(arg *UpdateFileMemberArgs) (res *MemberAccessLevelResult, err error) {
	return dbx.UpdateFileMemberContext(context.Background(), arg)
}

// UpdateFolderMemberAPIError is an error-wrapper for the update_folder_member route
type UpdateFolderMemberAPIError struct {
	dropbox.APIError
	EndpointError *UpdateFolderMemberError `json:"error"`
}

func (dbx *apiImpl) UpdateFolderMemberContext(ctx context.Context, arg *UpdateFolderMemberArg) (res *MemberAccessLevelResult, err error) {
	req := dropbox.Request{
		Host:         "api",
		Namespace:    "sharing",
		Route:        "update_folder_member",
		Auth:         "user",
		Style:        "rpc",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr UpdateFolderMemberAPIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return
	}

	_ = respBody
	return
}

func (dbx *apiImpl) UpdateFolderMember(arg *UpdateFolderMemberArg) (res *MemberAccessLevelResult, err error) {
	return dbx.UpdateFolderMemberContext(context.Background(), arg)
}

// UpdateFolderPolicyAPIError is an error-wrapper for the update_folder_policy route
type UpdateFolderPolicyAPIError struct {
	dropbox.APIError
	EndpointError *UpdateFolderPolicyError `json:"error"`
}

func (dbx *apiImpl) UpdateFolderPolicyContext(ctx context.Context, arg *UpdateFolderPolicyArg) (res *SharedFolderMetadata, err error) {
	req := dropbox.Request{
		Host:         "api",
		Namespace:    "sharing",
		Route:        "update_folder_policy",
		Auth:         "user",
		Style:        "rpc",
		Arg:          arg,
		ExtraHeaders: nil,
	}

	var resp []byte
	var respBody io.ReadCloser
	resp, respBody, err = (*dropbox.Context)(dbx).Execute(ctx, req, nil)
	if err != nil {
		var appErr UpdateFolderPolicyAPIError
		err = auth.ParseError(err, &appErr)
		if errors.Is(err, &appErr) {
			err = appErr
		}
		return
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return
	}

	_ = respBody
	return
}

func (dbx *apiImpl) UpdateFolderPolicy(arg *UpdateFolderPolicyArg) (res *SharedFolderMetadata, err error) {
	return dbx.UpdateFolderPolicyContext(context.Background(), arg)
}

// New returns a Client implementation for this namespace
func New(c dropbox.Config) Client {
	ctx := apiImpl(dropbox.NewContext(c))
	return &ctx
}
