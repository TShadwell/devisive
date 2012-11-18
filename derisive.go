package main
import(
	"log"
	"fmt"
	//"os/exec"
	"net/http"
	"github.com/danieldk/gosvm"
	"strconv"
	"encoding/json"
	"io/ioutil"
	"github.com/garyburd/go-mongo/mongo"
)
const(
	MODEL_PREFIX = "mod_"
)

var derisiveMongo derisiveMongoConn

type API struct{
	aLevelModels map[yOut]*gosvm.Model
}
type derisiveMongoConn struct{
	connection mongo.Conn
	database mongo.Database
	deprivation mongo.Collection
}
func connectToMongo() (mng derisiveMongoConn, err error){
	mng.connection, err = mongo.Dial("localhost")
	mng.database = mongo.Database{
		mng.connection,
		"derisive",
		nil,
	}
	mng.deprivation =mng.database.C("dep_mappings_2")
	return
}
func (dmc *derisiveMongoConn)getDeprivationData(lat, lng float32) ([]gosvm.FeatureValue, error){
	url:="http://mapit.mysociety.org/point/4326/" + fmt.Sprint(lng) + "," + fmt.Sprint(lat)+ "?type=OLF"
	res, err := http.Get(url)
	if err != nil {
		return make([]gosvm.FeatureValue, 0), err
	}
	jsn, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil{
		return make([]gosvm.FeatureValue, 0), err
	}
	var lowerLayerStruct interface{}
	err = json.Unmarshal(jsn, &lowerLayerStruct)
	if err != nil{
		return make([]gosvm.FeatureValue, 0), err
	}
	lls := lowerLayerStruct.(map[string]interface{})


	var ko string
	for ki, _:= range lls{
		ko = ki
	}
	obj := lls[ko].(map[string]interface{})

	codes := obj["codes"].(map[string]interface{})

	code := codes["ons"].(string)
	var deprdata map[string]interface{}
	dmc.deprivation.Find(mongo.M{
		"code":code,
	}).One(&deprdata)
	if err != nil && err != mongo.Done{
		log.Fatal(err)
	}
	_depdata := deprdata["dep"].([]interface{})
	/*
	for _, val := range _depdata{
		fmt.Println("Value", val.(string))
	}
	*/
	out := make([]gosvm.FeatureValue, len(_Deps))
	for i, val := range _Deps{
		magnitude, _ := strconv.ParseFloat(_depdata[i].(string), 64)
		out[i] = gosvm.FeatureValue{
			int(val),
			magnitude,
		}
	}
	return out, err
}
func (dmc *derisiveMongoConn) getSchoolData(schoolname string) 
func (ws *API) prepare(){
	//Make the map
	ws.aLevelModels = make(map[yOut]*gosvm.Model, NUM_YINP_VALUES)
	for i:=1;i< NUM_YINP_VALUES+1;i++{
		yO := yOut(i)
		var err error
		ws.aLevelModels[yO], err = gosvm.LoadModel(MODEL_PREFIX+strconv.Itoa(i))
		if err != nil {
			panic("Error loading model file for `" + yO.String() + "'.\n\t"+
				"Was expecting filename: " + MODEL_PREFIX+strconv.Itoa(i))
		}
		log.Println("Loaded module `" + yO.String() + "'.")
	}
}
func (ws *API) ServeHTTP(respond http.ResponseWriter, req *http.Request){
}
type web struct{

}
func (w *web) ServeHTTP(respond http.ResponseWriter, req *http.Request){

}
func main(){
	log.Println(NAME, " starting...")
	main := new(API)
	main.prepare()
	log.Println("Models loaded.")

	log.Println("Connecting to mongodb...")
	derisiveMongo, err := connectToMongo()
	if err != nil {
		log.Fatal("Could not connect to mongo!")
	}
	log.Println("\tConnected.")
	_, err = derisiveMongo.getDeprivationData(51.5133299,-0.088947)


	if err !=nil{
		log.Fatal(err)
	}

	log.Println("Starting HTTP server...")
	log.Println("\tStarted...")

}
