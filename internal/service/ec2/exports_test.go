// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

// Exports for use in tests only.
var (
	ResourceAMICopy                                       = resourceAMICopy
	ResourceAMIFromInstance                               = resourceAMIFromInstance
	ResourceAMILaunchPermission                           = resourceAMILaunchPermission
	ResourceAvailabilityZoneGroup                         = resourceAvailabilityZoneGroup
	ResourceCapacityReservation                           = resourceCapacityReservation
	ResourceCarrierGateway                                = resourceCarrierGateway
	ResourceClientVPNAuthorizationRule                    = resourceClientVPNAuthorizationRule
	ResourceClientVPNEndpoint                             = resourceClientVPNEndpoint
	ResourceClientVPNNetworkAssociation                   = resourceClientVPNNetworkAssociation
	ResourceClientVPNRoute                                = resourceClientVPNRoute
	ResourceCustomerGateway                               = resourceCustomerGateway
	ResourceDefaultNetworkACL                             = resourceDefaultNetworkACL
	ResourceDefaultRouteTable                             = resourceDefaultRouteTable
	ResourceEBSDefaultKMSKey                              = resourceEBSDefaultKMSKey
	ResourceEBSEncryptionByDefault                        = resourceEBSEncryptionByDefault
	ResourceEBSFastSnapshotRestore                        = newEBSFastSnapshotRestoreResource
	ResourceEBSSnapshot                                   = resourceEBSSnapshot
	ResourceEBSSnapshotCopy                               = resourceEBSSnapshotCopy
	ResourceEBSSnapshotImport                             = resourceEBSSnapshotImport
	ResourceEBSVolume                                     = resourceEBSVolume
	ResourceEIP                                           = resourceEIP
	ResourceEIPAssociation                                = resourceEIPAssociation
	ResourceEIPDomainName                                 = newEIPDomainNameResource
	ResourceFleet                                         = resourceFleet
	ResourceFlowLog                                       = resourceFlowLog
	ResourceHost                                          = resourceHost
	ResourceIPAM                                          = resourceIPAM
	ResourceIPAMOrganizationAdminAccount                  = resourceIPAMOrganizationAdminAccount
	ResourceIPAMPool                                      = resourceIPAMPool
	ResourceIPAMPoolCIDR                                  = resourceIPAMPoolCIDR
	ResourceIPAMPoolCIDRAllocation                        = resourceIPAMPoolCIDRAllocation
	ResourceIPAMPreviewNextCIDR                           = resourceIPAMPreviewNextCIDR
	ResourceIPAMResourceDiscovery                         = resourceIPAMResourceDiscovery
	ResourceIPAMResourceDiscoveryAssociation              = resourceIPAMResourceDiscoveryAssociation
	ResourceIPAMScope                                     = resourceIPAMScope
	ResourceImageBlockPublicAccess                        = resourceImageBlockPublicAccess
	ResourceInstance                                      = resourceInstance
	ResourceInstanceConnectEndpoint                       = newInstanceConnectEndpointResource
	ResourceInstanceMetadataDefaults                      = newInstanceMetadataDefaultsResource
	ResourceInstanceState                                 = resourceInstanceState
	ResourceInternetGateway                               = resourceInternetGateway
	ResourceInternetGatewayAttachment                     = resourceInternetGatewayAttachment
	ResourceKeyPair                                       = resourceKeyPair
	ResourceLaunchTemplate                                = resourceLaunchTemplate
	ResourceLocalGatewayRoute                             = resourceLocalGatewayRoute
	ResourceLocalGatewayRouteTableVPCAssociation          = resourceLocalGatewayRouteTableVPCAssociation
	ResourceMainRouteTableAssociation                     = resourceMainRouteTableAssociation
	ResourceManagedPrefixList                             = resourceManagedPrefixList
	ResourceManagedPrefixListEntry                        = resourceManagedPrefixListEntry
	ResourceNATGateway                                    = resourceNATGateway
	ResourceNetworkACL                                    = resourceNetworkACL
	ResourceNetworkACLAssociation                         = resourceNetworkACLAssociation
	ResourceNetworkACLRule                                = resourceNetworkACLRule
	ResourceNetworkInsightsAnalysis                       = resourceNetworkInsightsAnalysis
	ResourceNetworkInsightsPath                           = resourceNetworkInsightsPath
	ResourceNetworkInterface                              = resourceNetworkInterface
	ResourceNetworkInterfaceSGAttachment                  = resourceNetworkInterfaceSGAttachment
	ResourceNetworkPerformanceMetricSubscription          = resourceNetworkPerformanceMetricSubscription
	ResourcePlacementGroup                                = resourcePlacementGroup
	ResourceRoute                                         = resourceRoute
	ResourceRouteTable                                    = resourceRouteTable
	ResourceRouteTableAssociation                         = resourceRouteTableAssociation
	ResourceSecurityGroupEgressRule                       = newSecurityGroupEgressRuleResource
	ResourceSecurityGroupIngressRule                      = newSecurityGroupIngressRuleResource
	ResourceSnapshotCreateVolumePermission                = resourceSnapshotCreateVolumePermission
	ResourceSpotDataFeedSubscription                      = resourceSpotDataFeedSubscription
	ResourceSpotFleetRequest                              = resourceSpotFleetRequest
	ResourceSpotInstanceRequest                           = resourceSpotInstanceRequest
	ResourceSubnet                                        = resourceSubnet
	ResourceSubnetCIDRReservation                         = resourceSubnetCIDRReservation
	ResourceTag                                           = resourceTag
	ResourceTrafficMirrorFilter                           = resourceTrafficMirrorFilter
	ResourceTrafficMirrorFilterRule                       = resourceTrafficMirrorFilterRule
	ResourceTrafficMirrorSession                          = resourceTrafficMirrorSession
	ResourceTrafficMirrorTarget                           = resourceTrafficMirrorTarget
	ResourceTransitGatewayConnect                         = resourceTransitGatewayConnect
	ResourceTransitGatewayMulticastDomain                 = resourceTransitGatewayMulticastDomain
	ResourceTransitGatewayMulticastDomainAssociation      = resourceTransitGatewayMulticastDomainAssociation
	ResourceTransitGatewayMulticastGroupMember            = resourceTransitGatewayMulticastGroupMember
	ResourceTransitGatewayMulticastGroupSource            = resourceTransitGatewayMulticastGroupSource
	ResourceTransitGatewayPeeringAttachment               = resourceTransitGatewayPeeringAttachment
	ResourceTransitGatewayPeeringAttachmentAccepter       = resourceTransitGatewayPeeringAttachmentAccepter
	ResourceTransitGatewayPolicyTable                     = resourceTransitGatewayPolicyTable
	ResourceTransitGatewayPolicyTableAssociation          = resourceTransitGatewayPolicyTableAssociation
	ResourceTransitGatewayPrefixListReference             = resourceTransitGatewayPrefixListReference
	ResourceTransitGatewayRoute                           = resourceTransitGatewayRoute
	ResourceTransitGatewayRouteTable                      = resourceTransitGatewayRouteTable
	ResourceTransitGatewayRouteTableAssociation           = resourceTransitGatewayRouteTableAssociation
	ResourceTransitGatewayRouteTablePropagation           = resourceTransitGatewayRouteTablePropagation
	ResourceTransitGatewayVPCAttachment                   = resourceTransitGatewayVPCAttachment
	ResourceTransitGatewayVPCAttachmentAccepter           = resourceTransitGatewayVPCAttachmentAccepter
	ResourceVerifiedAccessEndpoint                        = resourceVerifiedAccessEndpoint
	ResourceVerifiedAccessGroup                           = resourceVerifiedAccessGroup
	ResourceVerifiedAccessInstance                        = resourceVerifiedAccessInstance
	ResourceVerifiedAccessInstanceLoggingConfiguration    = resourceVerifiedAccessInstanceLoggingConfiguration
	ResourceVerifiedAccessInstanceTrustProviderAttachment = resourceVerifiedAccessInstanceTrustProviderAttachment
	ResourceVerifiedAccessTrustProvider                   = resourceVerifiedAccessTrustProvider
	ResourceVPC                                           = resourceVPC
	ResourceVPCDHCPOptions                                = resourceVPCDHCPOptions
	ResourceVPCDHCPOptionsAssociation                     = resourceVPCDHCPOptionsAssociation
	ResourceVPCEndpoint                                   = resourceVPCEndpoint
	ResourceVPCEndpointPolicy                             = resourceVPCEndpointPolicy
	ResourceVPCEndpointRouteTableAssociation              = resourceVPCEndpointRouteTableAssociation
	ResourceVPCEndpointSecurityGroupAssociation           = resourceVPCEndpointSecurityGroupAssociation
	ResourceVPCEndpointService                            = resourceVPCEndpointService
	ResourceVPCEndpointSubnetAssociation                  = resourceVPCEndpointSubnetAssociation
	ResourceVPCIPv4CIDRBlockAssociation                   = resourceVPCIPv4CIDRBlockAssociation
	ResourceVPCPeeringConnection                          = resourceVPCPeeringConnection
	ResourceVPNConnection                                 = resourceVPNConnection
	ResourceVPNConnectionRoute                            = resourceVPNConnectionRoute
	ResourceVPNGateway                                    = resourceVPNGateway
	ResourceVPNGatewayAttachment                          = resourceVPNGatewayAttachment
	ResourceVPNGatewayRoutePropagation                    = resourceVPNGatewayRoutePropagation
	ResourceVolumeAttachment                              = resourceVolumeAttachment

	CustomFiltersSchema                                        = customFiltersSchema
	ErrCodeDefaultSubnetAlreadyExistsInAvailabilityZone        = errCodeDefaultSubnetAlreadyExistsInAvailabilityZone
	ErrCodeInvalidSpotDatafeedNotFound                         = errCodeInvalidSpotDatafeedNotFound
	ExpandIPPerms                                              = expandIPPerms
	FindAvailabilityZones                                      = findAvailabilityZones
	FindCapacityReservationByID                                = findCapacityReservationByID
	FindCarrierGatewayByID                                     = findCarrierGatewayByID
	FindClientVPNAuthorizationRuleByThreePartKey               = findClientVPNAuthorizationRuleByThreePartKey
	FindClientVPNEndpointByID                                  = findClientVPNEndpointByID
	FindClientVPNNetworkAssociationByTwoPartKey                = findClientVPNNetworkAssociationByTwoPartKey
	FindClientVPNRouteByThreePartKey                           = findClientVPNRouteByThreePartKey
	FindCreateSnapshotCreateVolumePermissionByTwoPartKey       = findCreateSnapshotCreateVolumePermissionByTwoPartKey
	FindCustomerGatewayByID                                    = findCustomerGatewayByID
	FindDHCPOptionsByID                                        = findDHCPOptionsByID
	FindEBSVolumeAttachment                                    = findVolumeAttachment
	FindEBSVolumeByID                                          = findEBSVolumeByID
	FindEgressOnlyInternetGatewayByID                          = findEgressOnlyInternetGatewayByID
	FindEIPByAllocationID                                      = findEIPByAllocationID
	FindEIPByAssociationID                                     = findEIPByAssociationID
	FindEIPDomainNameAttributeByAllocationID                   = findEIPDomainNameAttributeByAllocationID
	FindFastSnapshotRestoreByTwoPartKey                        = findFastSnapshotRestoreByTwoPartKey
	FindFleetByID                                              = findFleetByID
	FindFlowLogByID                                            = findFlowLogByID
	FindHostByID                                               = findHostByID
	FindInternetGateway                                        = findInternetGateway
	FindInternetGatewayAttachment                              = findInternetGatewayAttachment
	FindInternetGatewayByID                                    = findInternetGatewayByID
	FindIPAMByID                                               = findIPAMByID
	FindIPAMPoolAllocationByTwoPartKey                         = findIPAMPoolAllocationByTwoPartKey
	FindIPAMPoolByID                                           = findIPAMPoolByID
	FindIPAMPoolCIDRByTwoPartKey                               = findIPAMPoolCIDRByTwoPartKey
	FindIPAMResourceDiscoveryAssociationByID                   = findIPAMResourceDiscoveryAssociationByID
	FindIPAMResourceDiscoveryByID                              = findIPAMResourceDiscoveryByID
	FindIPAMScopeByID                                          = findIPAMScopeByID
	FindImageLaunchPermission                                  = findImageLaunchPermission
	FindInstanceConnectEndpointByID                            = findInstanceConnectEndpointByID
	FindInstanceMetadataDefaults                               = findInstanceMetadataDefaults
	FindInstanceStateByID                                      = findInstanceStateByID
	FindKeyPairByName                                          = findKeyPairByName
	FindLaunchTemplateByID                                     = findLaunchTemplateByID
	FindLocalGatewayRouteByTwoPartKey                          = findLocalGatewayRouteByTwoPartKey
	FindLocalGatewayRouteTableVPCAssociationByID               = findLocalGatewayRouteTableVPCAssociationByID
	FindMainRouteTableAssociationByID                          = findMainRouteTableAssociationByID
	FindManagedPrefixListByID                                  = findManagedPrefixListByID
	FindManagedPrefixListEntryByIDAndCIDR                      = findManagedPrefixListEntryByIDAndCIDR
	FindNATGatewayByID                                         = findNATGatewayByID
	FindNetworkACLAssociationByID                              = findNetworkACLAssociationByID
	FindNetworkACLByID                                         = findNetworkACLByID
	FindNetworkACLEntryByThreePartKey                          = findNetworkACLEntryByThreePartKey
	FindNetworkInsightsAnalysisByID                            = findNetworkInsightsAnalysisByID
	FindNetworkInsightsPathByID                                = findNetworkInsightsPathByID
	FindNetworkInterfaceByID                                   = findNetworkInterfaceByID
	FindNetworkInterfaceSecurityGroup                          = findNetworkInterfaceSecurityGroup
	FindNetworkPerformanceMetricSubscriptionByFourPartKey      = findNetworkPerformanceMetricSubscriptionByFourPartKey
	FindPlacementGroupByName                                   = findPlacementGroupByName
	FindPublicIPv4Pools                                        = findPublicIPv4Pools
	FindRouteByIPv4Destination                                 = findRouteByIPv4Destination
	FindRouteByIPv6Destination                                 = findRouteByIPv6Destination
	FindRouteByPrefixListIDDestination                         = findRouteByPrefixListIDDestination
	FindRouteTableAssociationByID                              = findRouteTableAssociationByID
	FindRouteTableByID                                         = findRouteTableByID
	FindSecurityGroupByID                                      = findSecurityGroupByID
	FindSecurityGroupEgressRuleByID                            = findSecurityGroupEgressRuleByID
	FindSecurityGroupIngressRuleByID                           = findSecurityGroupIngressRuleByID
	FindSnapshot                                               = findSnapshot
	FindSnapshotByID                                           = findSnapshotByID
	FindSpotDatafeedSubscription                               = findSpotDatafeedSubscription
	FindSpotFleetRequestByID                                   = findSpotFleetRequestByID
	FindSpotFleetRequests                                      = findSpotFleetRequests
	FindSpotInstanceRequestByID                                = findSpotInstanceRequestByID
	FindSubnetCIDRReservationBySubnetIDAndReservationID        = findSubnetCIDRReservationBySubnetIDAndReservationID
	FindSubnets                                                = findSubnets
	FindTag                                                    = findTag
	FindTrafficMirrorFilterByID                                = findTrafficMirrorFilterByID
	FindTrafficMirrorFilterRuleByTwoPartKey                    = findTrafficMirrorFilterRuleByTwoPartKey
	FindTrafficMirrorSessionByID                               = findTrafficMirrorSessionByID
	FindTrafficMirrorTargetByID                                = findTrafficMirrorTargetByID
	FindTransitGatewayByID                                     = findTransitGatewayByID
	FindTransitGatewayConnectByID                              = findTransitGatewayConnectByID
	FindTransitGatewayConnectPeerByID                          = findTransitGatewayConnectPeerByID
	FindTransitGatewayMulticastDomainAssociationByThreePartKey = findTransitGatewayMulticastDomainAssociationByThreePartKey
	FindTransitGatewayMulticastDomainByID                      = findTransitGatewayMulticastDomainByID
	FindTransitGatewayMulticastGroupMemberByThreePartKey       = findTransitGatewayMulticastGroupMemberByThreePartKey
	FindTransitGatewayMulticastGroupSourceByThreePartKey       = findTransitGatewayMulticastGroupSourceByThreePartKey
	FindTransitGatewayPeeringAttachmentByID                    = findTransitGatewayPeeringAttachmentByID
	FindTransitGatewayPolicyTableAssociationByTwoPartKey       = findTransitGatewayPolicyTableAssociationByTwoPartKey
	FindTransitGatewayPolicyTableByID                          = findTransitGatewayPolicyTableByID
	FindTransitGatewayPrefixListReferenceByTwoPartKey          = findTransitGatewayPrefixListReferenceByTwoPartKey
	FindTransitGatewayRouteTableAssociationByTwoPartKey        = findTransitGatewayRouteTableAssociationByTwoPartKey
	FindTransitGatewayRouteTableByID                           = findTransitGatewayRouteTableByID
	FindTransitGatewayRouteTablePropagationByTwoPartKey        = findTransitGatewayRouteTablePropagationByTwoPartKey
	FindTransitGatewayStaticRoute                              = findTransitGatewayStaticRoute
	FindTransitGatewayVPCAttachmentByID                        = findTransitGatewayVPCAttachmentByID
	FindVPCCIDRBlockAssociationByID                            = findVPCCIDRBlockAssociationByID
	FindVPCDHCPOptionsAssociation                              = findVPCDHCPOptionsAssociation
	FindVPCEndpointConnectionByServiceIDAndVPCEndpointID       = findVPCEndpointConnectionByServiceIDAndVPCEndpointID
	FindVPCEndpointConnectionNotificationByID                  = findVPCEndpointConnectionNotificationByID
	FindVPCEndpointRouteTableAssociationExists                 = findVPCEndpointRouteTableAssociationExists
	FindVPCEndpointSecurityGroupAssociationExists              = findVPCEndpointSecurityGroupAssociationExists
	FindVPCEndpointServiceConfigurationByID                    = findVPCEndpointServiceConfigurationByID
	FindVPCEndpointServicePermission                           = findVPCEndpointServicePermission
	FindVPCEndpointSubnetAssociationExists                     = findVPCEndpointSubnetAssociationExists
	FindVPCIPv6CIDRBlockAssociationByID                        = findVPCIPv6CIDRBlockAssociationByID
	FindVPCPeeringConnectionByID                               = findVPCPeeringConnectionByID
	FindVPNConnectionByID                                      = findVPNConnectionByID
	FindVPNConnectionRouteByTwoPartKey                         = findVPNConnectionRouteByTwoPartKey
	FindVPNGatewayByID                                         = findVPNGatewayByID
	FindVPNGatewayRoutePropagationExists                       = findVPNGatewayRoutePropagationExists
	FindVPNGatewayVPCAttachmentByTwoPartKey                    = findVPNGatewayVPCAttachmentByTwoPartKey
	FindVerifiedAccessEndpointByID                             = findVerifiedAccessEndpointByID
	FindVerifiedAccessGroupByID                                = findVerifiedAccessGroupByID
	FindVerifiedAccessInstanceByID                             = findVerifiedAccessInstanceByID
	FindVerifiedAccessInstanceLoggingConfigurationByInstanceID = findVerifiedAccessInstanceLoggingConfigurationByInstanceID
	FindVerifiedAccessInstanceTrustProviderAttachmentExists    = findVerifiedAccessInstanceTrustProviderAttachmentExists
	FindVerifiedAccessTrustProviderByID                        = findVerifiedAccessTrustProviderByID
	FindVolumeAttachmentInstanceByID                           = findVolumeAttachmentInstanceByID
	FlattenNetworkInterfacePrivateIPAddresses                  = flattenNetworkInterfacePrivateIPAddresses
	FlattenSecurityGroups                                      = flattenSecurityGroups
	InternetGatewayAttachmentParseResourceID                   = internetGatewayAttachmentParseResourceID
	IPAMServicePrincipal                                       = ipamServicePrincipal
	ManagedPrefixListEntryCreateResourceID                     = managedPrefixListEntryCreateResourceID
	ManagedPrefixListEntryParseResourceID                      = managedPrefixListEntryParseResourceID
	MatchRules                                                 = matchRules
	NewAttributeFilterList                                     = newAttributeFilterList
	NewAttributeFilterListV2                                   = newAttributeFilterListV2
	NewCustomFilterList                                        = newCustomFilterList
	NewTagFilterList                                           = newTagFilterList
	ProtocolForValue                                           = protocolForValue
	SecurityGroupExpandRules                                   = securityGroupExpandRules
	SecurityGroupCollapseRules                                 = securityGroupCollapseRules
	SecurityGroupIPPermGather                                  = securityGroupIPPermGather
	SecurityGroupMigrateState                                  = securityGroupMigrateState
	SecurityGroupRuleCreateID                                  = securityGroupRuleCreateID
	SecurityGroupRuleMigrateState                              = securityGroupRuleMigrateState
	SpotFleetRequestMigrateState                               = spotFleetRequestMigrateState
	StopEBSVolumeAttachmentInstance                            = stopVolumeAttachmentInstance
	StopInstance                                               = stopInstance
	UnsuccessfulItemError                                      = unsuccessfulItemError
	UnsuccessfulItemsError                                     = unsuccessfulItemsError
	UpdateTags                                                 = updateTags
	VPCDHCPOptionsAssociationParseResourceID                   = vpcDHCPOptionsAssociationParseResourceID
	VPCMigrateState                                            = vpcMigrateState
	WaitVolumeAttachmentCreated                                = waitVolumeAttachmentCreated
)

type (
	IPProtocol = ipProtocol
)
