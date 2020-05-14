package developMemory

const (
	developName = "develop"
	developPass = "develop"
)

var BasicDevelop Develop

type Develop struct {
	Name string
	Pass string
}

func SetDevPassword() {
	BasicDevelop = Develop{
		Name: developName,
		Pass: developPass,
	}
}

func UpdateDevPass(pass string) {
	BasicDevelop = Develop{
		Name: developName,
		Pass: pass,
	}
}
