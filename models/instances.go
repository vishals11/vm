package models

//aWS VM instance structure
type Instance struct {
	InstanceID       string `json:"instance_id,omitempty" gorm:"primary_key"`
	Name             string `json:"name,omitempty"`
	InstanceType     string `json:"instance_type,omitempty"`
	AvailabilityZone string `json:"availability_zone,omitempty"`
	InstanceState    string `json:"instance_state,omitempty"`
	StatusChecks     string `json:"status_checks,omitempty"`
	AlarmStatus      string `json:"alarm_status,omitempty"`
	PublicDNS        string `json:"public_dns,omitempty"`
	IPv4             string `json:"ipv4,omitempty"`
	IPv6             string `json:"ipv6,omitempty"`
	KeyName          string `json:"key_name,omitempty"`
	Monitoring       string `json:"monitoring,omitempty"`
	LaunchTime       string `json:"launch_time,omitempty"`
	SecurityGroups   string `json:"security_groups,omitempty"`
	Owner            string `json:"owner,omitempty"`
}

var instances []Instance

func CreateInstance(instance *Instance) (*Instance, error) {
	err := db.Create(&instance).Error
	if err != nil {
		return nil, err
	}
	return instance, nil
}

//Example : http://localhost:8000/v1/instances?pagesize=4&offset=0&order_by=instance_id desc
func GetAllInstances(pagesize int, offset int, order_by string) (*[]Instance, error) {
	var instances []Instance
	err := db.Limit(pagesize).Offset(offset).Order(order_by).Find(&instances).Error
	if err != nil {
		return nil, err
	}
	return &instances, nil
}

func GetInstance(instanceid *string, instance *Instance) (*Instance, error) {
	err := db.Where("instance_id=?", instanceid).Find(&instance).Error
	if err != nil {
		return nil, err
	}
	return instance, nil
}

func UpdateInstance(instanceid *string, instance *Instance) {
	db.Where("instance_id=?", instanceid).Find(&instance)
	db.Save(&instance)
}

func DeleteInstance(instanceid *string, instance *Instance) (*Instance, error) {
	err := db.Where("instance_id=?", instanceid).Find(&instance).Error
	if err != nil {
		return nil, err
	}
	db.Delete(&instance)
	return instance, nil
}
