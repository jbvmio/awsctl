package awsctl

// Instance holds metadata of an ec2 instance.
type Instance struct {
	AZ             string
	ID             string
	Image          string
	Index          int64
	KeyName        string
	Name           string
	PrivateDnsName string
	PrivateIP      string
	PublicDnsName  string
	PublicIP       string
	State          string
	Type           string
	VPC            string
	Tags           map[string]string
	TagCount       int
}

func (cl *Client) GetInstances(ids ...string) []Instance {
	return cl.GetInstanceMap().GetInstances(ids...)
}

func (i InstanceMap) GetInstances(ids ...string) []Instance {
	var instances []Instance
	switch {
	case len(ids) > 0:
		for _, id := range ids {
			if i[id] != nil {
				var name string
				var tags map[string]string
				if i[id].Tags != nil {
					tags = make(map[string]string, len(i[id].Tags))
					for _, tag := range i[id].Tags {
						tags[*tag.Key] = *tag.Value
					}
					name = tags["Name"]
				}
				inst := Instance{
					AZ:             *i[id].Placement.AvailabilityZone,
					ID:             *i[id].InstanceId,
					Image:          *i[id].ImageId,
					Index:          *i[id].AmiLaunchIndex,
					KeyName:        *i[id].KeyName,
					Name:           name,
					PrivateDnsName: *i[id].PrivateDnsName,
					PrivateIP:      *i[id].PrivateIpAddress,
					PublicDnsName:  *i[id].PublicDnsName,
					PublicIP:       *i[id].PublicIpAddress,
					State:          *i[id].State.Name,
					Type:           *i[id].InstanceType,
					VPC:            *i[id].VpcId,
					Tags:           tags,
					TagCount:       len(tags),
				}
				instances = append(instances, inst)
			}
		}
	default:
		for id := range i {
			var name string
			var tags map[string]string
			if i[id].Tags != nil {
				tags = make(map[string]string, len(i[id].Tags))
				for _, tag := range i[id].Tags {
					tags[*tag.Key] = *tag.Value
				}
				name = tags["Name"]
			}
			inst := Instance{
				AZ:             *i[id].Placement.AvailabilityZone,
				ID:             *i[id].InstanceId,
				Image:          *i[id].ImageId,
				Index:          *i[id].AmiLaunchIndex,
				KeyName:        *i[id].KeyName,
				Name:           name,
				PrivateDnsName: *i[id].PrivateDnsName,
				PrivateIP:      *i[id].PrivateIpAddress,
				PublicDnsName:  *i[id].PublicDnsName,
				PublicIP:       *i[id].PublicIpAddress,
				State:          *i[id].State.Name,
				Type:           *i[id].InstanceType,
				VPC:            *i[id].VpcId,
				Tags:           tags,
				TagCount:       len(tags),
			}
			instances = append(instances, inst)
		}
	}
	return instances
}
