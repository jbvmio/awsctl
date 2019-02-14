package awsctl

// Instance holds metadata of an ec2 instance.
type Instance struct {
	AZ          string
	ID          string
	Image       string
	KeyPair     string
	PrivateName string
	PrivateIP   string
	PublicName  string
	PublicIP    string
	State       string
	Type        string
	VPC         string
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
				inst := Instance{
					AZ:          *i[id].Placement.AvailabilityZone,
					ID:          *i[id].InstanceId,
					Image:       *i[id].ImageId,
					KeyPair:     *i[id].KeyName,
					PrivateName: *i[id].PrivateDnsName,
					PrivateIP:   *i[id].PrivateIpAddress,
					PublicName:  *i[id].PublicDnsName,
					PublicIP:    *i[id].PublicIpAddress,
					State:       *i[id].State.Name,
					Type:        *i[id].InstanceType,
					VPC:         *i[id].VpcId,
				}
				instances = append(instances, inst)
			}
		}
	default:
		for id := range i {
			inst := Instance{
				AZ:          *i[id].Placement.AvailabilityZone,
				ID:          *i[id].InstanceId,
				Image:       *i[id].ImageId,
				KeyPair:     *i[id].KeyName,
				PrivateName: *i[id].PrivateDnsName,
				PrivateIP:   *i[id].PrivateIpAddress,
				PublicName:  *i[id].PublicDnsName,
				PublicIP:    *i[id].PublicDnsName,
				State:       *i[id].State.Name,
				Type:        *i[id].InstanceType,
				VPC:         *i[id].VpcId,
			}
			instances = append(instances, inst)
		}
	}
	return instances
}
