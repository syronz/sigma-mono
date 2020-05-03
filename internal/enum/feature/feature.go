package feature

// Feature defines the type of feature
type Feature string

/* List of All features, different version has been created with variate
combination of these features

Activity: See all action on the application, by default it is saved, just showing it
disabled

DeleteList: Ability to see deleted rows and return back them in case of mistake

*/
const (
	Activity          Feature = "Activity"
	Autobackup        Feature = "Autobackup"
	BalanceSheet      Feature = "BalanceSheet"
	CashFlowStatement Feature = "CashFlowStatement"
	CurrencyNumber    Feature = "CurrencyNumber"
	DeleteList        Feature = "DeleteList"
	IncomeStatemnt    Feature = "IncomeStatemnt"
	Inventory         Feature = "Inventory"
	LogAPI            Feature = "LogAPI"
	MultiUser         Feature = "MultiUser"
	NetworkIntranet   Feature = "NetworkIntranet"
	NetworkLocal      Feature = "NetworkLocal"
	NetworkPublic     Feature = "NetworkPublic"
	Undelete          Feature = "Undelete"
)

// Features contain all type of Feature to be used in validation
var Features = []string{
	string(Activity),
	string(Autobackup),
	string(BalanceSheet),
	string(CashFlowStatement),
	string(CurrencyNumber),
	string(DeleteList),
	string(IncomeStatemnt),
	string(Inventory),
	string(LogAPI),
	string(MultiUser),
	string(NetworkIntranet),
	string(NetworkLocal),
	string(NetworkPublic),
	string(Undelete),
}
