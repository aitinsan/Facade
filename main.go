package main

import "fmt"

type person struct{

}
func (p *person) WriteStatement(){
	fmt.Println("Пишешь заявление")
}
func (p *person) ScanDocuments(){
	fmt.Println("Сканируешь все необходимые документы")
}
func (p* person) СonvictionFromEgov(){
	fmt.Println("Берешь справку что нет судимостой в egov")
}

type polyclinic struct{
}
func (p *polyclinic) Bloodanalyze(){
	fmt.Println("сдал кровь")
}
func (p *polyclinic) LORAnalyze(){
	fmt.Println("носовые полости рентгенили")
}
func (p *polyclinic) Fluro(){
	fmt.Println("сдал флюру")
}
type voenkomat struct {

}
func (v *voenkomat) Queue(){
	fmt.Println("Ждал огромную очередь")
}
func (v *voenkomat) CaseGet(){
	fmt.Println("забрал личное дело")
}
func (v *voenkomat) Lor(){
	fmt.Println("прошел лора")
}
func (v *voenkomat) Dentist(){
	fmt.Println("прошел стоматолга")
}
func (v *voenkomat) Oculist(){
	fmt.Println("прошел окулиста")
}
func (v *voenkomat) Dermatologist(){
	fmt.Println("прошел дерматолога")
}
func (v *voenkomat) Surgeon(){
	fmt.Println("прошел хирурга")
}
func (v *voenkomat) Sadness(){
	fmt.Println("Сказали что я не годен")
}
type kafedraFacade struct {
	person *person
	polyclinic *polyclinic
	voenkomat *voenkomat
}
func MyRoadToKafedra() *kafedraFacade{
	return &kafedraFacade{new(person),new(polyclinic),new(voenkomat)}
}
func (k *kafedraFacade) FirstActions(){
	k.person.WriteStatement()
	k.person.ScanDocuments()
	k.person.СonvictionFromEgov()

}
func (k *kafedraFacade) SecondActions(){
	k.polyclinic.Bloodanalyze()
	k.polyclinic.Fluro()
	k.polyclinic.LORAnalyze()
}
func (k *kafedraFacade) ThirdActions(){
	k.voenkomat.Queue()
	k.voenkomat.CaseGet()
	k.voenkomat.Dentist()
	k.voenkomat.Dermatologist()
	k.voenkomat.Lor()
	k.voenkomat.Oculist()
	k.voenkomat.Surgeon()
	k.voenkomat.Sadness()
}
func main(){
	me := MyRoadToKafedra()
	me.FirstActions()
	me.SecondActions()
	me.ThirdActions()

}