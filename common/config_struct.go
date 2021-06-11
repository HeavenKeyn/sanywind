package common

type FaultItem struct {
	VariableName string `json:"variable_name"`
	FaultCode    string `json:"fault_code"`
	FaultDesc    string `json:"fault_desc"`
	PartBelongs  string `json:"part_belongs"`
	AlarmLevel   string `json:"alarm_level"`
	BrakeProNO   string `json:"brake_no"`
}

type Farm struct {
	FarmCode     string  `json:"farm_code"`
	FarmName     string  `json:"farm_name"`
	Capacity     float64 `json:"capacity"`
	InstalledNum int     `json:"installed_num"`
	Owner        string  `json:"owner"`
	Country      string  `json:"country"`
	Province     string  `json:"province"`
	City         string  `json:"city"`
	Address      string  `json:"address"`
	Longitude    string  `json:"longitude"`
	Latitude     string  `json:"latitude"`
}

type Device struct {
	InnerTurbineName string `json:"inner_turbine_name"`
	InnerDeviceName  string `json:"inner_device_name"`
	OwnerDeviceName  string `json:"device_name"`
	PointVersion     string `json:"point_version"`
	PointName        string `json:"point_name"`
	FarmCode         string `json:"farm_code"`
	LoopName         string `json:"loop_name"`
	LoopOrder        int    `json:"loop_order"`
	Project          string `json:"project"`
}
type Turbine struct {
	InnerTurbineName  string  `json:"inner_turbine_name"`
	OwnerTurbineName  string  `json:"owner_turbine_name"`
	Type              string  `json:"type_name"`
	PointVersion      string  `json:"point_version"`
	PointName         string  `json:"point_name"`
	FarmCode          string  `json:"farm_code"`
	LoopName          string  `json:"loop_name"`
	LoopOrder         int     `json:"loop_order"`
	Project           string  `json:"project"`
	RatedPower        float64 `json:"rated_power"`
	RatedTorque       float64 `json:"rated_torque"`
	RatedSpeed        float64 `json:"rated_speed"`
	GridSpeed         float64 `json:"grid_speed"`
	CutInSpeed        float64 `json:"cut_in_speed"`
	CutOutSpeed       float64 `json:"cut_out_speed"`
	MinimumBladeAngle float64 `json:"minimum_blade_angle"`
	HubHeight         float64 `json:"hub_height"`
	Longitude         string  `json:"longitude"`
	Latitude          string  `json:"latitude"`
}
type AvailabilitySta struct {
	State string `json:"state"`
	Desc  string `json:"desc"`
}

type FirstFaultJson struct {
	StartTime    string `json:"startTime"`
	PointName    string `json:"point_name"`
	PointVersion string `json:"point_version"`
	VariableName string `json:"variable_name"`
	FaultCode    string `json:"fault_code"`
	FaultDesc    string `json:"fault_desc"`
	PartBelongs  string `json:"part_belongs"`
	AlarmLevel   string `json:"alarm_level"`
	BrakeProNO   string `json:"brake_no"`
}
type FaultVariable struct {
	VariableName string `json:"variable_name"`
}

type Point struct {
	RegisterType string `json:"register_type"`
	VariableName string `json:"variable_name"`
	DataType     string `json:"data_type"`
	Resample     string `json:"resample"`
}

type FarmDB struct {
	Farm   string              `json:"farm"`
	Port   string              `json:"port"`
	Tables map[string]TableUDP `json:"tables"`
}

type TableUDP struct {
	RetentionPolicy string `json:"retention-policy"`
	Precision       string `json:"precision"`
	Port            string `json:"port"`
}

type StatusReason struct {
	Code       string
	Des        string   `json:"des"`
	StatusCode []string `json:"status_code"`
	Judge      []string `json:"judge"`
}

type DeviceCtrlCommand struct {
	Agc     string `json:"agc"`
	Avc     string `json:"avc"`
	Stop    string `json:"stop"`
	Startup string `json:"startup"`
}

type ElectricityLoss struct {
	TheoreticalGeneration  float64 `db:"theoretical_generation"`
	WeatherLossGeneration  float64 `db:"weather_loss_generation"`
	OverhaulLossGeneration float64 `db:"overhaul_loss_generation"`
	MaintainLossGeneration float64 `db:"maintain_loss_generation"`
	FaultLossGeneration    float64 `db:"fault_loss_generation"`
	GridLossGeneration     float64 `db:"grid_loss_generation"`
	LocalLossGeneration    float64 `db:"local_loss_generation"`
	RemoteLossGeneration   float64 `db:"remote_loss_generation"`
	LimitLossGeneration    float64 `db:"limit_loss_generation"`
}
