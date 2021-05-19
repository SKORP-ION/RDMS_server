package database

/*
func (s SysinfoModel) Create(ws *structures.Workstation, sys *structures.Sysinfo) (SysinfoModel, error) {
	data, err := json.Marshal(sys)

	if err != nil {
		return s, errors.New("Can't parse JSON")
	}

	s.Info = string(data)

	if ws.Id == 0 {
		return s, errors.New("Workstation not found")
	}
	s.Id = ws.Id
	s.Ts = time.Now().UTC()

	return s, nil
}
*/

