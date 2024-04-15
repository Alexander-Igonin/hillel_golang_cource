package computer

type Computer struct {
	Cpu string
	Gpu string
	Motherboard string
	Ram string
	Ssd string
	PowerBlock string
	Water string
	
}

func NewComputer(cpu, gpu, motherboard, ram, ssd, powerblock, water string) *Computer {
	return &Computer{
		Cpu: cpu,
		Gpu: gpu,
		Motherboard: motherboard,
		Ram: ram,
		Ssd: ssd,
		PowerBlock: powerblock,
		Water: water,
	}
}
