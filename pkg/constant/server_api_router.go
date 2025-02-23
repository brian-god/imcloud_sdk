// Copyright © 2023 OpenIM SDK. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package constant

const (
	BaseRouter   = "/api/app/v1"
	BaseRouterV2 = "/api/app/v2"
)
const (
	GetSelfUserInfoRouter         = BaseRouter + "/user/get_pb_self_user_info"
	GetUsersInfoRouter            = BaseRouter + "/user/get_users_info"
	FindFullProfileByUserIdRouter = BaseRouter + "/user/find_full_users_info"
	UpdateSelfUserInfoRouter      = BaseRouter + "/user/update_user_info"
	SetGlobalRecvMessageOptRouter = BaseRouter + "/user/set_global_msg_recv_opt"
	GetUsersInfoFromCacheRouter   = BaseRouter + "/user/get_users_info_from_cache"
	SearchUserInfoRouter          = BaseRouter + "/user/search"
	SetUsersOption                = BaseRouter + "/user/set_option"
	GetUserLoginStatusRouter      = BaseRouter + "/user/get_login_status"
	GetUserOperation              = BaseRouter + "/user/get_user_operation"

	ScreenUserProfile = BaseRouter + "/user/screen_profile"

	AddFriendRouter                           = BaseRouter + "/friend/add_friend"
	DeleteFriendRouter                        = BaseRouter + "/friend/delete_friend"
	GetSelfFriendReceiveApplicationListRouter = BaseRouter + "/friend/get_self_receive_friend_apply_list" //recv
	GetSelfFriendApplicationListRouter        = BaseRouter + "/friend/get_self_friend_apply_list"         //send
	GetFriendListRouter                       = BaseRouter + "/friend/get_friend_list"
	AddFriendResponse                         = BaseRouter + "/friend/add_friend_response"
	SetFriendInfoRouter                       = BaseRouter + "/friend/set_friend_info"
	// 黑明单
	AddBlackRouter     = BaseRouter + "/friend/add_black"
	RemoveBlackRouter  = BaseRouter + "/friend/remove_black"
	GetBlackListRouter = BaseRouter + "/friend/get_black_list"
	// GetFriendRequestByApplicantRouter 通过申请与被申请人获取好友请求详情
	GetFriendRequestByApplicantRouter = BaseRouter + "/friend/get_friend_request_by_applicant"
	GetFriendByAppIdsRouter           = BaseRouter + "/friend/get_friend_by_ids"

	GetSyncFriendList      = BaseRouter + "/friend/get_sync_friend_list"
	SendMsgRouter          = "/chat/send_msg"
	PullUserMsgRouter      = "/chat/pull_msg"
	PullUserMsgBySeqRouter = BaseRouter + "/msg/pull_msg_by_seq"
	NewestSeqRouter        = "/chat/newest_seq"

	//msg
	ClearConversationMsgRouter             = BaseRouter + RouterMsg + "/clear_conversation_msg"
	ClearConversationsAllMsgRouter         = BaseRouter + RouterMsg + "/clear_conversation_all_msg" // Clear the message of the specified conversation
	ClearAllMsgRouter                      = BaseRouter + RouterMsg + "/user_clear_all_msg"         // Clear all messages of the current user
	DeleteMsgsRouter                       = BaseRouter + RouterMsg + "/delete_msgs"                // Delete the specified message
	RevokeMsgRouter                        = BaseRouter + RouterMsg + "/revoke_msg"
	SetMessageReactionExtensionsRouter     = BaseRouter + RouterMsg + "/set_message_reaction_extensions"
	AddMessageReactionExtensionsRouter     = BaseRouter + RouterMsg + "/add_message_reaction_extensions"
	MarkMsgsAsReadRouter                   = BaseRouter + RouterMsg + "/mark_msgs_as_read"
	GetConversationsHasReadAndMaxSeqRouter = BaseRouter + RouterMsg + "/get_conversations_has_read_and_max_seq"
	GetConversationMaxSeqRouter            = BaseRouter + RouterMsg + "/get_conversation_max_seq"

	MarkConversationAsRead    = BaseRouter + RouterMsg + "/mark_conversation_as_read"
	MarkMsgsAsRead            = BaseRouter + RouterMsg + "/mark_msgs_as_read"
	SetConversationHasReadSeq = BaseRouter + RouterMsg + "/set_conversation_has_read_seq"

	GetMessageListReactionExtensionsRouter = BaseRouter + RouterMsg + "/get_message_list_reaction_extensions"
	DeleteMessageReactionExtensionsRouter  = BaseRouter + RouterMsg + "/delete_message_reaction_extensions"

	TencentCloudStorageCredentialRouter = BaseRouter + "/third/tencent_cloud_storage_credential"
	AliOSSCredentialRouter              = BaseRouter + "/third/ali_oss_credential"
	MinioStorageCredentialRouter        = BaseRouter + "/third/minio_storage_credential"
	AwsStorageCredentialRouter          = BaseRouter + "/third/aws_storage_credential"

	//group
	CreateGroupRouter                 = BaseRouter + RouterGroup + "/create_group"
	SetGroupInfoRouter                = BaseRouter + RouterGroup + "/set_group_info"
	JoinGroupRouter                   = BaseRouter + RouterGroup + "/join_group"
	QuitGroupRouter                   = BaseRouter + RouterGroup + "/quit_group"
	GetGroupsInfoRouter               = BaseRouter + RouterGroup + "/get_groups_info"
	GetGroupMemberListRouter          = BaseRouter + RouterGroup + "/get_group_member_list"
	GetGroupAllMemberListRouter       = BaseRouter + RouterGroup + "/get_group_all_member_list"
	GetGroupMembersInfoRouter         = BaseRouter + RouterGroup + "/get_group_members_info"
	InviteUserToGroupRouter           = BaseRouter + RouterGroup + "/invite_user_to_group"
	GetJoinedGroupListRouter          = BaseRouter + RouterGroup + "/get_joined_group_list"
	KickGroupMemberRouter             = BaseRouter + RouterGroup + "/kick_group"
	TransferGroupRouter               = BaseRouter + RouterGroup + "/transfer_group"
	GetRecvGroupApplicationListRouter = BaseRouter + RouterGroup + "/get_recv_group_applicationList"
	// 以群主或管理员身份获取未处理的加群请求
	GetUntreatedRecvGroupApplicationListRouter = BaseRouter + RouterGroup + "/get_untreated_recv_group_application_list"

	FindFullGroupInfoRouter           = BaseRouter + RouterGroup + "/find_full_group_info"
	GetSendGroupApplicationListRouter = BaseRouter + RouterGroup + "/get_user_req_group_applicationList"
	AcceptGroupApplicationRouter      = BaseRouter + RouterGroup + "/group_application_response"
	RefuseGroupApplicationRouter      = BaseRouter + RouterGroup + "/group_application_response"
	DismissGroupRouter                = BaseRouter + RouterGroup + "/dismiss_group"
	MuteGroupMemberRouter             = BaseRouter + RouterGroup + "/mute_group_member"
	CancelMuteGroupMemberRouter       = BaseRouter + RouterGroup + "/cancel_mute_group_member"
	MuteGroupRouter                   = BaseRouter + RouterGroup + "/mute_group"
	CancelMuteGroupRouter             = BaseRouter + RouterGroup + "/cancel_mute_group"
	SetGroupMemberNicknameRouter      = BaseRouter + RouterGroup + "/set_group_member_nickname"
	SetGroupMemberInfoRouter          = BaseRouter + RouterGroup + "/set_group_member_info"
	GetGroupAbstractInfoRouter        = BaseRouter + RouterGroup + "/get_group_abstract_info"
	SearchGroupInfoRouter             = BaseRouter + RouterGroup + "/search_group_info"
	SearchGroupByCodeRouter           = BaseRouter + RouterGroup + "/search_group_by_code"
	GetJoinGroupRequestDetailRouter   = BaseRouter + RouterGroup + "/get_join_group_request_detail"
	GetUserOwnerJoinRequestNumRouter  = BaseRouter + RouterGroup + "/get_user_owner_join_request_num"
	// GetGroupMemberByIdsRouter 根据群成员id获取群信息
	GetGroupMemberByIdsRouter = BaseRouter + RouterGroup + "/get_member_by_ids"

	GetSyncGroupInfoList = BaseRouter + RouterGroup + "/get_sync_group_list"

	UpdateGroupSwitch = BaseRouter + RouterGroup + "/set_group_switch_info"
	//同步群成员
	SyncGroupMemberInfoRouter = BaseRouter + RouterGroup + "/sync_group_members"
	//根据时间同步群
	SyncUserJoinGroupInfoByTimeRouter = BaseRouter + RouterGroup + "/sync_group_info_by_time"
	//根据时间同步群成员
	SyncGroupMemberInfoUpdateTimeRouter = BaseRouter + RouterGroup + "/sync_group_member_info_by_time"

	GetUserMemberInfoInGroup = BaseRouter + RouterGroup + "/get_member_in_group"
	group

	SetReceiveMessageOptRouter         = BaseRouter + "/conversation/set_receive_message_opt"
	GetReceiveMessageOptRouter         = BaseRouter + "/conversation/get_receive_message_opt"
	GetAllConversationMessageOptRouter = BaseRouter + "/conversation/get_all_conversation_message_opt"
	SetConversationOptRouter           = BaseRouter + ConversationGroup + "/set_conversation"
	GetConversationsRouter             = BaseRouter + ConversationGroup + "/get_conversations"
	GetAllConversationsRouter          = BaseRouter + ConversationGroup + "/get_all_conversations"
	GetConversationRouter              = BaseRouter + ConversationGroup + "/get_conversation"
	BatchSetConversationRouter         = BaseRouter + ConversationGroup + "/batch_set_conversation"
	ModifyConversationFieldRouter      = BaseRouter + ConversationGroup + "/modify_conversation_field"
	SetConversationsRouter             = BaseRouter + ConversationGroup + "/set_conversations"

	//organization
	GetSubDepartmentRouter    = BaseRouter + RouterOrganization + "/get_sub_department"
	GetDepartmentMemberRouter = BaseRouter + RouterOrganization + "/get_department_member"
	ParseTokenRouter          = BaseRouter + RouterAuth + "/parse_token"

	//super_group
	GetJoinedSuperGroupListRouter = BaseRouter + RouterSuperGroup + "/get_joined_group_list"
	GetSuperGroupsInfoRouter      = BaseRouter + RouterSuperGroup + "/get_groups_info"

	//third
	FcmUpdateTokenRouter = BaseRouter + RouterThird + "/fcm_update_token"
	SetAppBadgeRouter    = BaseRouter + RouterThird + "/set_app_badge"
)
const (
	RouterGroup        = "/group"
	ConversationGroup  = "/conversation"
	RouterOrganization = "/organization"
	RouterAuth         = "/auth"
	RouterSuperGroup   = "/super_group"
	RouterMsg          = "/msg"
	RouterThird        = "/third"
	Moments            = "/moments"
)
const (
	ObjectPartLimit               = BaseRouter + "/third/part_limit"
	ObjectPartSize                = BaseRouter + "/third/part_size"
	ObjectInitiateMultipartUpload = BaseRouter + "/third/init_multi_upload"
	ObjectAuthSign                = BaseRouter + "/third/auth_sign"
	ObjectCompleteMultipartUpload = BaseRouter + "/third/complete_multi_upload"
	ObjectAccessURL               = BaseRouter + "/third/access_url"
)

const (
	V2ListMomentsRouter    = BaseRouterV2 + Moments + "/list"
	V2PublishMomentsRouter = BaseRouterV2 + Moments + "/publish"
	V2CommentMomentsRouter = BaseRouterV2 + Moments + "/comment"
	V2DeleteMomentsRouter  = BaseRouterV2 + Moments + "/delete"
	V2LikeMomentsRouter    = BaseRouterV2 + Moments + "/like"
	V2UnlikeMomentsRouter  = BaseRouterV2 + Moments + "/unlike"
)

//const (
//	GetSelfUserInfoRouter         = "/user/get_self_user_info"
//	GetUsersInfoRouter            = "/user/get_users_info"
//	UpdateSelfUserInfoRouter      = "/user/update_user_info"
//	UpdateSelfUserInfoExRouter    = "/user/update_user_info_ex"
//	SetGlobalRecvMessageOptRouter = "/user/set_global_msg_recv_opt"
//	ProcessUserCommandAdd         = "/user/process_user_command_add"
//	ProcessUserCommandDelete      = "/user/process_user_command_delete"
//	ProcessUserCommandUpdate      = "/user/process_user_command_update"
//	ProcessUserCommandGet         = "/user/process_user_command_get"
//	ProcessUserCommandGetAll      = "/user/process_user_command_get_all"
//
//	GetUsersInfoFromCacheRouter   = "/user/get_users_info_from_cache"
//	AccountCheck                  = "/user/account_check"
//	UserRegister                  = "/user/user_register"
//	SubscribeUsersStatusRouter    = "/user/subscribe_users_status"
//	GetSubscribeUsersStatusRouter = "/user/get_subscribe_users_status"
//	GetUserStatusRouter           = "/user/get_users_status"
//
//	AddFriendRouter                    = "/friend/add_friend"
//	DeleteFriendRouter                 = "/friend/delete_friend"
//	GetFriendApplicationListRouter     = "/friend/get_friend_apply_list"      //recv
//	GetSelfFriendApplicationListRouter = "/friend/get_self_friend_apply_list" //send
//	ImportFriendListRouter             = "/friend/import_friend"
//
//	GetDesignatedFriendsApplyRouter = "/friend/get_designated_friend_apply"
//	GetFriendListRouter             = "/friend/get_friend_list"
//	GetDesignatedFriendsRouter      = "/friend/get_designated_friends"
//	AddFriendResponse               = "/friend/add_friend_response"
//	SetFriendRemark                 = "/friend/set_friend_remark"
//	UpdateFriends                   = "/friend/update_friends"
//
//	AddBlackRouter     = "/friend/add_black"
//	RemoveBlackRouter  = "/friend/remove_black"
//	GetBlackListRouter = "/friend/get_black_list"
//
//	PullUserMsgRouter      = "/chat/pull_msg"
//	PullUserMsgBySeqRouter = "/chat/pull_msg_by_seq"
//	NewestSeqRouter        = "/chat/newest_seq"
//
//	// msg
//	ClearConversationMsgRouter             = RouterMsg + "/clear_conversation_msg" // Clear the message of the specified conversation
//	ClearAllMsgRouter                      = RouterMsg + "/user_clear_all_msg"     // Clear all messages of the current user
//	DeleteMsgsRouter                       = RouterMsg + "/delete_msgs"            // Delete the specified message
//	RevokeMsgRouter                        = RouterMsg + "/revoke_msg"
//	SetMessageReactionExtensionsRouter     = RouterMsg + "/set_message_reaction_extensions"
//	AddMessageReactionExtensionsRouter     = RouterMsg + "/add_message_reaction_extensions"
//	MarkMsgsAsReadRouter                   = RouterMsg + "/mark_msgs_as_read"
//	GetConversationsHasReadAndMaxSeqRouter = RouterMsg + "/get_conversations_has_read_and_max_seq"
//
//	MarkConversationAsRead    = RouterMsg + "/mark_conversation_as_read"
//	MarkMsgsAsRead            = RouterMsg + "/mark_msgs_as_read"
//	SetConversationHasReadSeq = RouterMsg + "/set_conversation_has_read_seq"
//	SendMsgRouter             = RouterMsg + "/send_msg"
//	GetServerTimeRouter       = RouterMsg + "/get_server_time"
//
//	GetMessageListReactionExtensionsRouter = RouterMsg + "/get_message_list_reaction_extensions"
//	DeleteMessageReactionExtensionsRouter  = RouterMsg + "/delete_message_reaction_extensions"
//
//	TencentCloudStorageCredentialRouter = "/third/tencent_cloud_storage_credential"
//	AliOSSCredentialRouter              = "/third/ali_oss_credential"
//	MinioStorageCredentialRouter        = "/third/minio_storage_credential"
//	AwsStorageCredentialRouter          = "/third/aws_storage_credential"
//
//	// group
//	CreateGroupRouter                 = RouterGroup + "/create_group"
//	SetGroupInfoRouter                = RouterGroup + "/set_group_info"
//	JoinGroupRouter                   = RouterGroup + "/join_group"
//	QuitGroupRouter                   = RouterGroup + "/quit_group"
//	GetGroupsInfoRouter               = RouterGroup + "/get_groups_info"
//	GetGroupMemberListRouter          = RouterGroup + "/get_group_member_list"
//	GetGroupAllMemberListRouter       = RouterGroup + "/get_group_all_member_list"
//	GetGroupMembersInfoRouter         = RouterGroup + "/get_group_members_info"
//	InviteUserToGroupRouter           = RouterGroup + "/invite_user_to_group"
//	GetJoinedGroupListRouter          = RouterGroup + "/get_joined_group_list"
//	KickGroupMemberRouter             = RouterGroup + "/kick_group"
//	TransferGroupRouter               = RouterGroup + "/transfer_group"
//	GetRecvGroupApplicationListRouter = RouterGroup + "/get_recv_group_applicationList"
//	GetSendGroupApplicationListRouter = RouterGroup + "/get_user_req_group_applicationList"
//	AcceptGroupApplicationRouter      = RouterGroup + "/group_application_response"
//	RefuseGroupApplicationRouter      = RouterGroup + "/group_application_response"
//	DismissGroupRouter                = RouterGroup + "/dismiss_group"
//	MuteGroupMemberRouter             = RouterGroup + "/mute_group_member"
//	CancelMuteGroupMemberRouter       = RouterGroup + "/cancel_mute_group_member"
//	MuteGroupRouter                   = RouterGroup + "/mute_group"
//	CancelMuteGroupRouter             = RouterGroup + "/cancel_mute_group"
//	SetGroupMemberNicknameRouter      = RouterGroup + "/set_group_member_nickname"
//	SetGroupMemberInfoRouter          = RouterGroup + "/set_group_member_info"
//	GetGroupAbstractInfoRouter        = RouterGroup + "/get_group_abstract_info"
//
//	SetReceiveMessageOptRouter         = "/conversation/set_receive_message_opt"
//	GetReceiveMessageOptRouter         = "/conversation/get_receive_message_opt"
//	GetAllConversationMessageOptRouter = "/conversation/get_all_conversation_message_opt"
//	SetConversationOptRouter           = ConversationGroup + "/set_conversation"
//	GetConversationsRouter             = ConversationGroup + "/get_conversations"
//	GetAllConversationsRouter          = ConversationGroup + "/get_all_conversations"
//	GetConversationRouter              = ConversationGroup + "/get_conversation"
//	BatchSetConversationRouter         = ConversationGroup + "/batch_set_conversation"
//	ModifyConversationFieldRouter      = ConversationGroup + "/modify_conversation_field"
//	SetConversationsRouter             = ConversationGroup + "/set_conversations"
//
//	// organization
//	GetSubDepartmentRouter    = RouterOrganization + "/get_sub_department"
//	GetDepartmentMemberRouter = RouterOrganization + "/get_department_member"
//	ParseTokenRouter          = RouterAuth + "/parse_token"
//
//	// super_group
//	GetJoinedSuperGroupListRouter = RouterSuperGroup + "/get_joined_group_list"
//	GetSuperGroupsInfoRouter      = RouterSuperGroup + "/get_groups_info"
//
//	// third
//	FcmUpdateTokenRouter = RouterThird + "/fcm_update_token"
//	SetAppBadgeRouter    = RouterThird + "/set_app_badge"
//	UploadLogsRouter     = RouterThird + "/logs/upload"
//
//	// auth
//	GetUsersToken = RouterAuth + "/user_token"
//)
//const (
//	RouterGroup        = "/group"
//	ConversationGroup  = "/conversation"
//	RouterOrganization = "/organization"
//	RouterAuth         = "/auth"
//	RouterSuperGroup   = "/super_group"
//	RouterMsg          = "/msg"
//	RouterThird        = "/third"
//)
//
//const (
//	ObjectPartLimit               = "/object/part_limit"
//	ObjectPartSize                = "/object/part_size"
//	ObjectInitiateMultipartUpload = "/object/initiate_multipart_upload"
//	ObjectAuthSign                = "/object/auth_sign"
//	ObjectCompleteMultipartUpload = "/object/complete_multipart_upload"
//	ObjectAccessURL               = "/object/access_url"
//)
