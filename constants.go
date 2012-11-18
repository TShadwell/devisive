package main
import(
	"strconv"
)
const(
	NUM_CORES	= 2
	NAME		= "Derisive"
)
//inputs
type xInp uint16
const (
	SCH_SCHOOLNAME xInp = iota+1
	SCH_ADDRESS1
	SCH_ADDRESS2
	SCH_ADDRESS3
	SCH_TOWN
	SCH_POSTCODE
	SCH_NFTYPE
	CEN_FSMEligible
	CEN_EthnicGroupMinor
	CEN_SENProvision
	CEN_LSOA
	CEN_AgeAtStartOfAcademicYear
	CEN_LanguageGroupMinor
	KS2_ENGLEV
	KS2_MATLEV
	KS2_SCILEV
	KS4_LEVEL2_EM
	KS4_PASS_AASTAR
	KS4_PASS_AC
	KS4_PASS_AG
	KS4_EBACC
	KS4_PTSTNEWG
	KS4_AVPTSENT
	KS4_APAASCI
	KS4_APADTSCI
	KS4_APAPDSCI
	KS4_APARA
	KS4_APART
	KS4_APBENa6
	KS4_APBHEB
	KS4_APBIO
	KS4_APBIOE
	KS4_APBUS
	KS4_APCGRK
	KS4_APCHE
	KS4_APCHI
	KS4_APCORESCI
	KS4_APDAN
	KS4_APDRA
	KS4_APDSCI
	KS4_APDTT
	KS4_APDUT
	KS4_APELEC
	KS4_APELIT
	KS4_APENG
	KS4_APFINE
	KS4_APFOOD
	KS4_APFRE
	KS4_APGEO
	KS4_APGER
	KS4_APGRA
	KS4_APGUJ
	KS4_APHECD
	KS4_APHIN
	KS4_APHIS
	KS4_APHSC
	KS4_APIT
	KS4_APITA
	KS4_APITSC
	KS4_APJAP
	KS4_APLAT
	KS4_APLT
	KS4_APMAT
	KS4_APMFT
	KS4_APMGRK
	KS4_APMHEB
	KS4_APMUS
	KS4_APOFT
	KS4_APPAN
	KS4_APPE
	KS4_APPER
	KS4_APPHY
	KS4_APPOL
	KS4_APPOR
	KS4_APRE
	KS4_APRES
	KS4_APRS
	KS4_APRUS
	KS4_APSPAN
	KS4_APSSC
	KS4_APSSCI
	KS4_APSTAT
	KS4_APSVSCI
	KS4_APSYS
	KS4_APTUR
	KS4_APURD
	KS4_APVBUS
	KS4_APVIT
	KS4_APVSCI
	KS4_APVSL
	KS4_BTECAPPSCI_CERS
	KS4_BTECAPPSCI_DIPS
	KS4_BTECENG_CERS
	KS4_BTECENG_DIPS
	KS5_ESTAB
	KS5_GENDER
	KS5_LA
	KS5_QROUTE
	KS5_QROUTEGS
	KS5_ToE_CODE
	KS4_APELANG
	//Index of Multiple Deprivation Score
	DEP1
	//Rank of Index of Multiple Deprivation Score
	DEP2
	//Income Score
	DEP3
	//Rank of Income Score
	DEP4
	//Employment Score
	DEP5
	//Rank of Employment Score
	DEP6
	//Health Deprivation and Disability Score
	DEP7
	//Rank of Health Deprivation and Disability Score
	DEP8
	//Education Skills and Training Score
	DEP9
	//Rank of Education Skills and Training Score
	DEP10
	//Barriers to Housing and Services Score
	DEP11
	//Rank of Barriers to Housing and Services Score
	DEP12
	//Crime Score
	DEP13
	//Rank of Crime Score
	DEP14
	//Living Environment Score
	DEP15
	//Rank of Living Environment Score
	DEP16

	LURN
	LLA
	LESTAB
	LLAESTAB
	LSCHNAME
	LSTREET
	LLOCALITY
	LADDRESS3
	LTOWN
	LPOSTCODE
	LTELNUM
	LICLOSE
	LISNEW
	LMINORGROUP
	LNFTYPE
	LISPRIMARY
	LISSECONDARY
	LISPOST16
	LAGEL
	LAGEH
	LGENDER
	LSFGENDER
	LRELDENOM
	LADMPOL
	LNEWACFLAG
	KS5_11RECTYPE
	KS5_11ALPHAIND
	KS5_11REGION
	KS5_11LASORT
	KS5_11LEA
	KS5_11ESTAB
	KS5_11URN
	KS5_11SCHNAME_AC
	KS5_11SCHNAME
	KS5_11ADDRESS1
	KS5_11ADDRESS2
	KS5_11ADDRESS3
	KS5_11TOWN
	KS5_11PCODE
	KS5_11TELNUM
	KS5_11CONTFLAG
	KS5_11NFTYPE
	KS5_11RELDENOM
	KS5_11ADMPOL
	KS5_11GENDER1618
	KS5_11FEEDER
	KS5_11AGERANGE
	KS5_11ICLOSE
	KS5_11TABKS2
	KS5_11TAB15
	KS5_11EXAMCONF
	KS5_11DUMMY1
	KS5_11TPUP1618
	KS5_11TALLPUPA
	KS5_11TALLPPSA
	KS5_11TALLPPEA
	KS5_11PTPASS1L3
	KS5_11PTPASS2LV3
	KS5_11PTPASS3LV3
	KS5_11TALLPPS08
	KS5_11TALLPPS09
	KS5_11TALLPPS10
	KS5_11TALLPPE08
	KS5_11TALLPPE09
	KS5_11TALLPPE10
	ABS_11LA
	ABS_11ESTAB
	ABS_11URN
	ABS_11PERCTOT
	ABS_11PERCUA
	ABS_11PPERSABS15
	ABS_11PPERSABS20
	CFR_11URN
	CFR_11LANUMBER
	//CFR_11LONDON/NON-LONDON
	LONDON
	CFR_11MEDIAN
	CFR_11PUPILS
	CFR_11FSM
	CFR_11FSMBAND
	CFR_11GRANTFUNDING
	CFR_11SELFGENINCOME
	CFR_11TOTALINCOME
	CFR_11TEACHINGSTAFF
	CFR_11SUPPLYTEACHERS
	CFR_11EDUCATIONSUPPORTSTAFF
	CFR_11PREMISES
	CFR_11BACKOFFICE
	CFR_11CATERING
	CFR_11OTHERSTAFF
	CFR_11ENERGY
	CFR_11LEARNINGRESOURCES
	CFR_11ICT
	CFR_11BOUGHTIN
	CFR_11OTHER
	CFR_11TOTALEXPENDITURE
	SWF_11LA
	SWF_11URN
	SWF_11NTEA
	SWF_11NTEAAS
	SWF_11NNONTEA
	SWF_11NFTETEA
	SWF_11NFTETEAAS
	SWF_11RATPUPTEA
	SWF_11SALARY
	CENSUS_11URN
	CENSUS_11LAESTAB
	CENSUS_11NUMFTE
	CENSUS_11TOTPUPSENDN
	CENSUS_11TSENSAP
	CENSUS_11TSENA
	CENSUS_11TOTSENST
	CENSUS_11TOTSENAP
	CENSUS_11PSENSAP
	CENSUS_11PSENA
	CENSUS_11PTOTSENST
	CENSUS_11PTOTSENAP
	CENSUS_11TOTPUPEALDN
	CENSUS_11NUMEAL
	CENSUS_11NUMENGFL
	CENSUS_11NUMUNCFL
	CENSUS_11PNUMEAL
	CENSUS_11PNUMENGFL
	CENSUS_11PNUMUNCFL
	CENSUS_11TOTPUPFSMDN
	CENSUS_11NUMFSM
	CENSUS_11NUMNOFSM
	CENSUS_11PNUMFSM
	CENSUS_11PNUMNOFSM
	OLA
	OURN
	OSCHOOLNAME
	OPHASE
	OREPORTURL
	//School Index of Multiple Deprivation Score
	SCHOOL1
	//School Rank of Index of Multiple Deprivation Score
	SCHOOL2
	//School Income Score
	SCHOOL3
	//School Rank of Income Score
	SCHOOL4
	//School Employment Score
	SCHOOL5
	//School Rank of Employment Score
	SCHOOL6
	//School Health Deprivation and Disability Score
	SCHOOL7
	//School Rank of Health Deprivation and Disability Score
	SCHOOL8
	//School Education Skills and Training Score
	SCHOOL9
	//School Rank of Education Skills and Training Score
	SCHOOL10
	//School Barriers to Housing and Services Score
	SCHOOL11
	//School Rank of Barriers to Housing and Services Score
	SCHOOL12
	//School Crime Score
	SCHOOL13
	//School Rank of Crime Score
	SCHOOL14
	//School Living Environment Score
	SCHOOL15
	//School Rank of Living Environment Score
	SCHOOL16


	//Thanks to go's wonderful iota operator
	NUM_XINP_VALUES = iota -1
)
var _Deps[]xInp = []xInp{
	DEP1,
	DEP2,
	DEP3,
	DEP4,
	DEP5,
	DEP6,
	DEP7,
	DEP8,
	DEP9,
	DEP10,
	DEP11,
	DEP12,
	DEP13,
	DEP14,
	DEP15,
	DEP16,
}
type yOut uint16
const (
	ks5_GA_ACCOUNTING yOut = iota +1
	ks5_GA_AD
	ks5_GA_AD_CRITI
	ks5_GA_AD_GRAPH
	ks5_GA_AD_PHOTO
	ks5_GA_AD_TEXTI
	ks5_GA_AD_THREE
	ks5_GA_ANC_HIST
	ks5_GA_ARABIC
	ks5_GA_ARCHAE
	ks5_GA_BENGALI
	ks5_GA_BIOLOGY
	ks5_GA_BIOLOGY_HUMAN
	ks5_GA_BUS
	ks5_GA_BUS_ECON
	ks5_GA_CHEMISTRY
	ks5_GA_CHINESE
	ks5_GA_CLASS_CIV
	ks5_GA_COMMUNICATION
	ks5_GA_COMP_STU
	ks5_GA_CRIT_THINK
	ks5_GA_DANCE
	ks5_GA_DRAMA
	ks5_GA_DT_FOOD
	ks5_GA_DT_PRODUCTION
	ks5_GA_DT_SYSTEMS
	ks5_GA_DUTCH
	ks5_GA_ECON
	ks5_GA_ELECTRONICS
	ks5_GA_ENG
	ks5_GA_ENG_LANG
	ks5_GA_ENG_LIT
	ks5_GA_ENV_SCI
	ks5_GA_EUROPE
	ks5_GA_FILM
	ks5_GA_FINE_ART
	ks5_GA_FRENCH
	ks5_GA_GEN_STUD
	ks5_GA_GEOG
	ks5_GA_GEOLOGY
	ks5_GA_GERMAN
	ks5_GA_GOV_POLITICS
	ks5_GA_GREEK
	ks5_GA_GUJARATI
	ks5_GA_HIST
	ks5_GA_HIST_ART
	ks5_GA_HOME_EC
	ks5_GA_IT
	ks5_GA_ITALIAN
	ks5_GA_JAPANESE
	ks5_GA_LATIN
	ks5_GA_LAW
	ks5_GA_LOGIC_PHIL
	ks5_GA_MATH
	ks5_GA_MATH_ADDI
	ks5_GA_MATH_APPL
	ks5_GA_MATH_DISC
	ks5_GA_MATH_FURT
	ks5_GA_MATH_MECH
	ks5_GA_MATH_PURE
	ks5_GA_MATH_STAT
	ks5_GA_MEDIA_FILM_TV
	ks5_GA_MOD_GREEK
	ks5_GA_MOD_HEBREW
	ks5_GA_MUSIC
	ks5_GA_MUSIC_TECH
	ks5_GA_OTH_CLASS
	ks5_GA_PANJABI
	ks5_GA_PE
	ks5_GA_PERFORMING
	ks5_GA_PERSIAN
	ks5_GA_PHYSICS
	ks5_GA_POLISH
	ks5_GA_PORTUGUESE
	ks5_GA_PSYCH_SCI
	ks5_GA_PSYCH_SOC
	ks5_GA_RE
	ks5_GA_RUSSIAN
	ks5_GA_SCI_PUBLIC
	ks5_GA_SCIENCE
	ks5_GA_SCO_POL
	ks5_GA_SCO_SCI_CIT
	ks5_GA_SOC
	ks5_GA_SPANISH
	ks5_GA_TURKISH
	ks5_GA_URDU
	ks5_GA_WELSH_SECOND
	ks5_GA_WORLD_DEV
	ks5_GAS_ACCOUNTING
	ks5_GAS_AD
	ks5_GAS_AD_CRITI
	ks5_GAS_AD_GRAPH
	ks5_GAS_AD_PHOTO
	ks5_GAS_AD_TEXTI
	ks5_GAS_AD_THREE
	ks5_GAS_ANC_HIST
	ks5_GAS_ARABIC
	ks5_GAS_ARCHAE
	ks5_GAS_BENGALI
	ks5_GAS_BIOLOGY
	ks5_GAS_BIOLOGY_HUMAN
	ks5_GAS_BUS
	ks5_GAS_BUS_ECON
	ks5_GAS_CHEMISTRY
	ks5_GAS_CHINESE
	ks5_GAS_CLASS_CIV
	ks5_GAS_COMMUNICATION
	ks5_GAS_COMP_STU
	ks5_GAS_CRIT_THINK
	ks5_GAS_DANCE
	ks5_GAS_DRAMA
	ks5_GAS_DT_FOOD
	ks5_GAS_DT_PRODUCTION
	ks5_GAS_DT_SYSTEMS
	ks5_GAS_DUTCH
	ks5_GAS_ECON
	ks5_GAS_ELECTRONICS
	ks5_GAS_ENG
	ks5_GAS_ENG_LANG
	ks5_GAS_ENG_LIT
	ks5_GAS_ENV_SCI
	ks5_GAS_EUROPE
	ks5_GAS_FILM
	ks5_GAS_FINE_ART
	ks5_GAS_FRENCH
	ks5_GAS_GEN_STUD
	ks5_GAS_GEOG
	ks5_GAS_GEOLOGY
	ks5_GAS_GERMAN
	ks5_GAS_GOV_POLITICS
	ks5_GAS_GREEK
	ks5_GAS_GUJARATI
	ks5_GAS_HIST
	ks5_GAS_HIST_ART
	ks5_GAS_HOME_EC
	ks5_GAS_IT
	ks5_GAS_ITALIAN
	ks5_GAS_JAPANESE
	ks5_GAS_LATIN
	ks5_GAS_LAW
	ks5_GAS_LOGIC_PHIL
	ks5_GAS_MATH
	ks5_GAS_MATH_ADDI
	ks5_GAS_MATH_APPL
	ks5_GAS_MATH_DISC
	ks5_GAS_MATH_FURT
	ks5_GAS_MATH_MECH
	ks5_GAS_MATH_PURE
	ks5_GAS_MATH_STAT
	ks5_GAS_MEDIA_FILM_TV
	ks5_GAS_MOD_GREEK
	ks5_GAS_MOD_HEBREW
	ks5_GAS_MUSIC
	ks5_GAS_MUSIC_TECH
	ks5_GAS_OTH_CLASS
	ks5_GAS_PANJABI
	ks5_GAS_PE
	ks5_GAS_PERFORMING
	ks5_GAS_PERSIAN
	ks5_GAS_PHYSICS
	ks5_GAS_POLISH
	ks5_GAS_PORTUGUESE
	ks5_GAS_PSYCH_SCI
	ks5_GAS_PSYCH_SOC
	ks5_GAS_RE
	ks5_GAS_RUSSIAN
	ks5_GAS_SCI_PUBLIC
	ks5_GAS_SCIENCE
	ks5_GAS_SCO_POL
	ks5_GAS_SCO_SCI_CIT
	ks5_GAS_SOC
	ks5_GAS_SPANISH
	ks5_GAS_TURKISH
	ks5_GAS_URDU
	ks5_GAS_WELSH_SECOND
	ks5_GAS_WORLD_DEV
	ks5_GDAS_AD
	ks5_GDAS_BUS
	ks5_GDAS_HEAL_SOC
	ks5_GDAS_ICT
	ks5_GDAS_LEIS_RECR
	ks5_GDAS_SCIENCE
	ks5_GDAS_TRAV_TOUR
	NUM_YINP_VALUES = iota -1
)
const (
)
func (y yOut) String()string{
	switch y +1 {
		case ks5_GA_ACCOUNTING: return `A Level Accounting/Finance`
		case ks5_GA_AD: return `A Level Art and Design`
		case ks5_GA_AD_CRITI: return `A Level Art and Design (Critical Studies)`
		case ks5_GA_AD_GRAPH: return `A Level Art and Design (Graphics)`
		case ks5_GA_AD_PHOTO: return `A Level Art and Design (Photography)`
		case ks5_GA_AD_TEXTI: return `A Level Art and Design (Textiles)`
		case ks5_GA_AD_THREE: return `A Level Art and Design (3-D Studies)`
		case ks5_GA_ANC_HIST: return `A Level Ancient History`
		case ks5_GA_ARABIC: return `A Level Arabic`
		case ks5_GA_ARCHAE: return `A Level Archaeology`
		case ks5_GA_BENGALI: return `A Level Bengali`
		case ks5_GA_BIOLOGY: return `A Level Biology`
		case ks5_GA_BIOLOGY_HUMAN: return `A Level Biology: Human`
		case ks5_GA_BUS: return `A Level Business Studies`
		case ks5_GA_BUS_ECON: return `A Level Business Studies and Economics`
		case ks5_GA_CHEMISTRY: return `A Level Chemistry`
		case ks5_GA_CHINESE: return `A Level Chinese`
		case ks5_GA_CLASS_CIV: return `A Level Classical Civilisation`
		case ks5_GA_COMMUNICATION: return `A Level Communication Studies`
		case ks5_GA_COMP_STU: return `A Level Computer Studies/Computing`
		case ks5_GA_CRIT_THINK: return `A Level Critical Thinking`
		case ks5_GA_DANCE: return `A Level Dance`
		case ks5_GA_DRAMA: return `A Level Drama & Theatre Studies`
		case ks5_GA_DT_FOOD: return `A Level Design/Tech & Food Technology`
		case ks5_GA_DT_PRODUCTION: return `A Level Design/Tech & Production Design`
		case ks5_GA_DT_SYSTEMS: return `A Level Design/Tech & Systems`
		case ks5_GA_DUTCH: return `A Level Dutch`
		case ks5_GA_ECON: return `A Level Economics`
		case ks5_GA_ELECTRONICS: return `A Level Science: Electronics`
		case ks5_GA_ENG: return `A Level English`
		case ks5_GA_ENG_LANG: return `A Level English Language`
		case ks5_GA_ENG_LIT: return `A Level English Literature`
		case ks5_GA_ENV_SCI: return `A Level Science: Environmental`
		case ks5_GA_EUROPE: return `A Level European Studies`
		case ks5_GA_FILM: return `A Level Film Studies`
		case ks5_GA_FINE_ART: return `A Level Fine Art`
		case ks5_GA_FRENCH: return `A Level French`
		case ks5_GA_GEN_STUD: return `A Level General Studies`
		case ks5_GA_GEOG: return `A Level Geography`
		case ks5_GA_GEOLOGY: return `A Level Science: Geology`
		case ks5_GA_GERMAN: return `A Level German`
		case ks5_GA_GOV_POLITICS: return `A Level Government & Politics`
		case ks5_GA_GREEK: return `A Level Greek`
		case ks5_GA_GUJARATI: return `A Level Gujarati`
		case ks5_GA_HIST: return `A Level History`
		case ks5_GA_HIST_ART: return `A Level History of Art`
		case ks5_GA_HOME_EC: return `A Level Home Economics`
		case ks5_GA_IT: return `A Level Information Technology`
		case ks5_GA_ITALIAN: return `A Level Italian`
		case ks5_GA_JAPANESE: return `A Level Japanese`
		case ks5_GA_LATIN: return `A Level Latin`
		case ks5_GA_LAW: return `A Level Law`
		case ks5_GA_LOGIC_PHIL: return `A Level Logic/Philosophy`
		case ks5_GA_MATH: return `A Level Mathematics`
		case ks5_GA_MATH_ADDI: return `A Level Additional Mathematics`
		case ks5_GA_MATH_APPL: return `A Level Mathematics (Applied)`
		case ks5_GA_MATH_DISC: return `A Level Mathematics (Discrete)`
		case ks5_GA_MATH_FURT: return `A Level Mathematics (Further)`
		case ks5_GA_MATH_MECH: return `A Level Mathematics (Mechanics)`
		case ks5_GA_MATH_PURE: return `A Level Mathematics (Pure)`
		case ks5_GA_MATH_STAT: return `A Level Mathematics (Statistics)`
		case ks5_GA_MEDIA_FILM_TV: return `A Level Media/Film/Television Studies`
		case ks5_GA_MOD_GREEK: return `A Level Modern Greek`
		case ks5_GA_MOD_HEBREW: return `A Level Modern Hebrew`
		case ks5_GA_MUSIC: return `A Level Music`
		case ks5_GA_MUSIC_TECH: return `A Level Music Technology`
		case ks5_GA_OTH_CLASS: return `A Level Other Classical Languages`
		case ks5_GA_PANJABI: return `A Level Punjabi`
		case ks5_GA_PE: return `A Level Sport/Physical Education Studies`
		case ks5_GA_PERFORMING: return `A Level Performing Studies`
		case ks5_GA_PERSIAN: return `A Level Persian`
		case ks5_GA_PHYSICS: return `A Level Physics`
		case ks5_GA_POLISH: return `A Level Polish`
		case ks5_GA_PORTUGUESE: return `A Level Portuguese`
		case ks5_GA_PSYCH_SCI: return `A Level Psychology JMB/NEA`
		case ks5_GA_PSYCH_SOC: return `A Level Psychology`
		case ks5_GA_RE: return `A Level Religious Studies`
		case ks5_GA_RUSSIAN: return `A Level Russian`
		case ks5_GA_SCI_PUBLIC: return `A Level Science for Public Understanding`
		case ks5_GA_SCIENCE: return `A Level Science: Single award`
		case ks5_GA_SCO_POL: return `A Level Social Policy`
		case ks5_GA_SCO_SCI_CIT: return `A Level Social Science Citizenship`
		case ks5_GA_SOC: return `A Level Sociology`
		case ks5_GA_SPANISH: return `A Level Spanish`
		case ks5_GA_TURKISH: return `A Level Turkish`
		case ks5_GA_URDU: return `A Level Urdu`
		case ks5_GA_WELSH_SECOND: return `A Level Welsh Second Language`
		case ks5_GA_WORLD_DEV: return `A Level World Development`
		case ks5_GAS_ACCOUNTING: return `AS Level Accounting/Finance`
		case ks5_GAS_AD: return `AS Level Art and Design`
		case ks5_GAS_AD_CRITI: return `AS Level Art and Design (Critical Studies)`
		case ks5_GAS_AD_GRAPH: return `AS Level Art and Design (Graphics)`
		case ks5_GAS_AD_PHOTO: return `AS Level Art and Design (Photography)`
		case ks5_GAS_AD_TEXTI: return `AS Level Art and Design (Textiles)`
		case ks5_GAS_AD_THREE: return `AS Level Art and Design (3-D Studies)`
		case ks5_GAS_ANC_HIST: return `AS Level Ancient History`
		case ks5_GAS_ARABIC: return `AS Level Arabic`
		case ks5_GAS_ARCHAE: return `AS Level Archaeology`
		case ks5_GAS_BENGALI: return `AS Level Bengali`
		case ks5_GAS_BIOLOGY: return `AS Level Biology`
		case ks5_GAS_BIOLOGY_HUMAN: return `AS Level Biology: Human`
		case ks5_GAS_BUS: return `AS Level Business Studies`
		case ks5_GAS_BUS_ECON: return `AS Level Business Studies and Economics`
		case ks5_GAS_CHEMISTRY: return `AS Level Chemistry`
		case ks5_GAS_CHINESE: return `AS Level Chinese`
		case ks5_GAS_CLASS_CIV: return `AS Level Classical Civilisation`
		case ks5_GAS_COMMUNICATION: return `AS Level Communication Studies`
		case ks5_GAS_COMP_STU: return `AS Level Computer Studies/Computing`
		case ks5_GAS_CRIT_THINK: return `AS Level Critical Thinking`
		case ks5_GAS_DANCE: return `AS Level Dance`
		case ks5_GAS_DRAMA: return `AS Level Drama & Theatre Studies`
		case ks5_GAS_DT_FOOD: return `AS Level Design/Tech & Food Technology`
		case ks5_GAS_DT_PRODUCTION: return `AS Level Design/Tech & Production Design`
		case ks5_GAS_DT_SYSTEMS: return `AS Level Design/Tech & Systems`
		case ks5_GAS_DUTCH: return `AS Level Dutch`
		case ks5_GAS_ECON: return `AS Level Economics`
		case ks5_GAS_ELECTRONICS: return `AS Level Science: Electronics`
		case ks5_GAS_ENG: return `AS Level English`
		case ks5_GAS_ENG_LANG: return `AS Level English Language`
		case ks5_GAS_ENG_LIT: return `AS Level English Literature`
		case ks5_GAS_ENV_SCI: return `AS Level Science: Environmental`
		case ks5_GAS_EUROPE: return `AS Level European Studies`
		case ks5_GAS_FILM: return `AS Level Film Studies`
		case ks5_GAS_FINE_ART: return `AS Level Fine Art`
		case ks5_GAS_FRENCH: return `AS Level French`
		case ks5_GAS_GEN_STUD: return `AS Level General Studies`
		case ks5_GAS_GEOG: return `AS Level Geography`
		case ks5_GAS_GEOLOGY: return `AS Level Science: Geology`
		case ks5_GAS_GERMAN: return `AS Level German`
		case ks5_GAS_GOV_POLITICS: return `AS Level Government & Politics`
		case ks5_GAS_GREEK: return `AS Level Greek`
		case ks5_GAS_GUJARATI: return `AS Level Gujarati`
		case ks5_GAS_HIST: return `AS Level History`
		case ks5_GAS_HIST_ART: return `AS Level History of Art`
		case ks5_GAS_HOME_EC: return `AS Level Home Economics`
		case ks5_GAS_IT: return `AS Level Information Technology`
		case ks5_GAS_ITALIAN: return `AS Level Italian`
		case ks5_GAS_JAPANESE: return `AS Level Japanese`
		case ks5_GAS_LATIN: return `AS Level Latin`
		case ks5_GAS_LAW: return `AS Level Law`
		case ks5_GAS_LOGIC_PHIL: return `AS Level Logic/Philosophy`
		case ks5_GAS_MATH: return `AS Level Mathematics`
		case ks5_GAS_MATH_ADDI: return `AS Level Additional Mathematics`
		case ks5_GAS_MATH_APPL: return `AS Level Mathematics (Applied)`
		case ks5_GAS_MATH_DISC: return `AS Level Mathematics (Discrete)`
		case ks5_GAS_MATH_FURT: return `AS Level Mathematics (Further)`
		case ks5_GAS_MATH_MECH: return `AS Level Mathematics (Mechanics)`
		case ks5_GAS_MATH_PURE: return `AS Level Mathematics (Pure)`
		case ks5_GAS_MATH_STAT: return `AS Level Mathematics (Statistics)`
		case ks5_GAS_MEDIA_FILM_TV: return `AS Level Media/Film/Television Studies`
		case ks5_GAS_MOD_GREEK: return `AS Level Modern Greek`
		case ks5_GAS_MOD_HEBREW: return `AS Level Modern Hebrew`
		case ks5_GAS_MUSIC: return `AS Level Music`
		case ks5_GAS_MUSIC_TECH: return `AS Level Music Technology`
		case ks5_GAS_OTH_CLASS: return `AS Level Other Classical Languages`
		case ks5_GAS_PANJABI: return `AS Level Punjabi`
		case ks5_GAS_PE: return `AS Level Sport/Physical Education Studies`
		case ks5_GAS_PERFORMING: return `AS Level Performing Studies`
		case ks5_GAS_PERSIAN: return `AS Level Persian`
		case ks5_GAS_PHYSICS: return `AS Level Physics`
		case ks5_GAS_POLISH: return `AS Level Polish`
		case ks5_GAS_PORTUGUESE: return `AS Level Portuguese`
		case ks5_GAS_PSYCH_SCI: return `AS Level Psychology JMB/NEA`
		case ks5_GAS_PSYCH_SOC: return `AS Level Psychology`
		case ks5_GAS_RE: return `AS Level Religious Studies`
		case ks5_GAS_RUSSIAN: return `AS Level Russian`
		case ks5_GAS_SCI_PUBLIC: return `AS Level Science for Public Understanding`
		case ks5_GAS_SCIENCE: return `AS Level Science: Single award`
		case ks5_GAS_SCO_POL: return `AS Level Social Policy`
		case ks5_GAS_SCO_SCI_CIT: return `AS Level Social Science Citizenship`
		case ks5_GAS_SOC: return `AS Level Sociology`
		case ks5_GAS_SPANISH: return `AS Level Spanish`
		case ks5_GAS_TURKISH: return `AS Level Turkish`
		case ks5_GAS_URDU: return `AS Level Urdu`
		case ks5_GAS_WELSH_SECOND: return `AS Level Welsh Second Language`
		case ks5_GAS_WORLD_DEV: return `AS Level World Development`
		case ks5_GDAS_AD: return `AS Double Award Level Applied Art and Design`
		case ks5_GDAS_BUS: return `AS Double Award Level Applied Business`
		case ks5_GDAS_HEAL_SOC: return `AS Double Award Level Health and Social Care`
		case ks5_GDAS_ICT: return `AS Double Award Level Applied ICT`
		case ks5_GDAS_LEIS_RECR: return `AS Double Award Level Leisure and Recreation`
		case ks5_GDAS_SCIENCE: return `AS Double Award Level Applied Science`
		case ks5_GDAS_TRAV_TOUR: return `AS Double Award Level Travel and Tourism`
	}
	panic("Output type '"+strconv.Itoa(int(y))+"' does not exist.")
}
