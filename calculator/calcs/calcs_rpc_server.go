package calcs

//CalcsrpcServer has to be implemented by the plugin
type CalcsrpcServer struct {
	// This is the real implementation
	Impl Calcs
}

//Calcs has to be implemented by the server which is the plugin
func (s *CalcsrpcServer) Calcs(args []float32, resp *float32) error {
	//result := fmt.Sprintf("%v",args)
	*resp = s.Impl.Operation(args)
	return nil
}
