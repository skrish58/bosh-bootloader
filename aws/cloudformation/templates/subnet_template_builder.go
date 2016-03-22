package templates

import "fmt"

type SubnetTemplateBuilder struct {
}

func NewSubnetTemplateBuilder() SubnetTemplateBuilder {
	return SubnetTemplateBuilder{}
}

func (s SubnetTemplateBuilder) BOSHSubnet() Template {
	return Template{
		Parameters: map[string]Parameter{
			"BOSHSubnetCIDR": Parameter{
				Description: "CIDR block for the BOSH subnet.",
				Type:        "String",
				Default:     "10.0.0.0/24",
			},
		},
		Resources: map[string]Resource{
			"BOSHSubnet": Resource{
				Type: "AWS::EC2::Subnet",
				Properties: Subnet{
					VpcId:     Ref{"VPC"},
					CidrBlock: Ref{"BOSHSubnetCIDR"},
					Tags: []Tag{
						{
							Key:   "Name",
							Value: "BOSH",
						},
					},
				},
			},
			"BOSHRouteTable": Resource{
				Type: "AWS::EC2::RouteTable",
				Properties: RouteTable{
					VpcId: Ref{"VPC"},
				},
			},
			"BOSHRoute": Resource{
				Type:      "AWS::EC2::Route",
				DependsOn: "VPCGatewayInternetGateway",
				Properties: Route{
					DestinationCidrBlock: "0.0.0.0/0",
					GatewayId:            Ref{"VPCGatewayInternetGateway"},
					RouteTableId:         Ref{"BOSHRouteTable"},
				},
			},
			"BOSHSubnetRouteTableAssociation": Resource{
				Type: "AWS::EC2::SubnetRouteTableAssociation",
				Properties: SubnetRouteTableAssociation{
					RouteTableId: Ref{"BOSHRouteTable"},
					SubnetId:     Ref{"BOSHSubnet"},
				},
			},
		},
		Outputs: map[string]Output{
			"BOSHSubnet": Output{
				Value: Ref{"BOSHSubnet"},
			},
			"BOSHSubnetAZ": Output{
				Value: FnGetAtt{
					[]string{
						"BOSHSubnet",
						"AvailabilityZone",
					},
				},
			},
		},
	}
}

func (s SubnetTemplateBuilder) InternalSubnet(azIndex int, suffix, cidrBlock string) Template {
	subnetName := fmt.Sprintf("InternalSubnet%s", suffix)
	subnetTag := fmt.Sprintf("Internal%s", suffix)
	subnetCIDRName := fmt.Sprintf("%sCIDR", subnetName)
	cidrDescription := fmt.Sprintf("CIDR block for %s.", subnetName)
	subnetRouteTableAssociationName := fmt.Sprintf("%sRouteTableAssociation", subnetName)
	return Template{
		Outputs: map[string]Output{
			fmt.Sprintf("%sName", subnetName): Output{
				Value: Ref{subnetName},
			},
			fmt.Sprintf("%sAZ", subnetName): Output{
				FnGetAtt{
					[]string{
						subnetName,
						"AvailabilityZone",
					},
				},
			},
			subnetCIDRName: Output{
				Value: Ref{subnetCIDRName},
			},
			fmt.Sprintf("%sSecurityGroup", subnetName): Output{
				Value: Ref{"InternalSecurityGroup"},
			},
		},
		Parameters: map[string]Parameter{
			subnetCIDRName: Parameter{
				Description: cidrDescription,
				Type:        "String",
				Default:     cidrBlock,
			},
		},
		Resources: map[string]Resource{
			subnetName: Resource{
				Type: "AWS::EC2::Subnet",
				Properties: Subnet{
					AvailabilityZone: map[string]interface{}{
						"Fn::Select": []interface{}{
							fmt.Sprintf("%d", azIndex),
							map[string]Ref{
								"Fn::GetAZs": Ref{"AWS::Region"},
							},
						},
					},
					CidrBlock: Ref{subnetCIDRName},
					VpcId:     Ref{"VPC"},
					Tags: []Tag{
						{
							Key:   "Name",
							Value: subnetTag,
						},
					},
				},
			},
			"InternalRouteTable": {
				Type: "AWS::EC2::RouteTable",
				Properties: RouteTable{
					VpcId: Ref{"VPC"},
				},
			},
			"InternalRoute": {
				Type:      "AWS::EC2::Route",
				DependsOn: "NATInstance",
				Properties: Route{
					DestinationCidrBlock: "0.0.0.0/0",
					RouteTableId:         Ref{"InternalRouteTable"},
					InstanceId:           Ref{"NATInstance"},
				},
			},
			subnetRouteTableAssociationName: Resource{
				Type: "AWS::EC2::SubnetRouteTableAssociation",
				Properties: SubnetRouteTableAssociation{
					RouteTableId: Ref{"InternalRouteTable"},
					SubnetId:     Ref{subnetName},
				},
			},
		},
	}
}

func (s SubnetTemplateBuilder) LoadBalancerSubnet() Template {
	return Template{
		Parameters: map[string]Parameter{
			"LoadBalancerSubnetCIDR": Parameter{
				Description: "CIDR block for the ELB subnet.",
				Type:        "String",
				Default:     "10.0.2.0/24",
			},
		},
		Resources: map[string]Resource{
			"LoadBalancerSubnet": Resource{
				Type: "AWS::EC2::Subnet",
				Properties: Subnet{
					AvailabilityZone: map[string]interface{}{
						"Fn::Select": []interface{}{
							"0",
							map[string]Ref{
								"Fn::GetAZs": Ref{"AWS::Region"},
							},
						},
					},
					CidrBlock: Ref{"LoadBalancerSubnetCIDR"},
					VpcId:     Ref{"VPC"},
					Tags: []Tag{
						{
							Key:   "Name",
							Value: "LoadBalancer",
						},
					},
				},
			},
			"LoadBalancerRouteTable": Resource{
				Type: "AWS::EC2::RouteTable",
				Properties: RouteTable{
					VpcId: Ref{"VPC"},
				},
			},
			"LoadBalancerRoute": Resource{
				Type:      "AWS::EC2::Route",
				DependsOn: "VPCGatewayInternetGateway",
				Properties: Route{
					DestinationCidrBlock: "0.0.0.0/0",
					GatewayId:            Ref{"VPCGatewayInternetGateway"},
					RouteTableId:         Ref{"LoadBalancerRouteTable"},
				},
			},
			"LoadBalancerSubnetRouteTableAssociation": {
				Type: "AWS::EC2::SubnetRouteTableAssociation",
				Properties: SubnetRouteTableAssociation{
					RouteTableId: Ref{"LoadBalancerRouteTable"},
					SubnetId:     Ref{"LoadBalancerSubnet"},
				},
			},
		},
	}
}
