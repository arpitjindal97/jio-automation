package main

import (
	"github.com/therecipe/qt/core"
)

func init() { CustomTableModel_QmlRegisterType2("CustomQmlTypes", 1, 0, "CustomTableModel") }

const (
	IMEI = int(core.Qt__UserRole) + 1<<iota
	Distributor
	Retailer
	Head1
	Date
)

type TableItem struct {
	imei        string
	distributor string
	retailer    string
	head1       string
	date        string
}

type CustomTableModel struct {
	core.QAbstractTableModel

	_ func() `constructor:"init"`

	_ func()                                  `signal:"remove,auto"`
	_ func(item TableItem)                    `signal:"add,auto"`
	_ func(firstName string, lastName string) `signal:"edit,auto"`
	_ func(filePath string)                   `signal:"openfile,auto"`

	modelData []TableItem
}

func (m *CustomTableModel) init() {
	m.modelData = tabledata
	m.ConnectRoleNames(m.roleNames)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectData(m.data)
}

func (m *CustomTableModel) roleNames() map[int]*core.QByteArray {
	return map[int]*core.QByteArray{
		IMEI:        core.NewQByteArray2("IMEI", -1),
		Distributor: core.NewQByteArray2("Distributor", -1),
		Retailer:    core.NewQByteArray2("Retailer", -1),
		Head1:       core.NewQByteArray2("Head1", -1),
		Date:        core.NewQByteArray2("Date", -1),
	}
}

func (m *CustomTableModel) rowCount(*core.QModelIndex) int {
	return len(m.modelData)
}

func (m *CustomTableModel) columnCount(*core.QModelIndex) int {
	return 5
}

func (m *CustomTableModel) data(index *core.QModelIndex, role int) *core.QVariant {
	item := m.modelData[index.Row()]
	switch role {
	case IMEI:
		return core.NewQVariant14(item.imei)
	case Distributor:
		return core.NewQVariant14(item.distributor)
	case Retailer:
		return core.NewQVariant14(item.retailer)
	case Head1:
		return core.NewQVariant14(item.head1)
	case Date:
		return core.NewQVariant14(item.date)

	}
	return core.NewQVariant()
}

func (m *CustomTableModel) remove() {
	if len(m.modelData) == 0 {
		return
	}
	m.BeginRemoveRows(core.NewQModelIndex(), len(m.modelData)-1, len(m.modelData)-1)
	m.modelData = m.modelData[:len(m.modelData)-1]
	m.EndRemoveRows()
}

func (m *CustomTableModel) add(item TableItem) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.modelData), len(m.modelData))
	m.modelData = append(m.modelData, item)
	m.EndInsertRows()
}

func (m *CustomTableModel) edit(firstName string, lastName string) {
	if len(m.modelData) == 0 {
		return
	}
	//m.modelData[len(m.modelData)-1] = TableItem{firstName, lastName}
	//m.DataChanged(m.Index(len(m.modelData)-1, 0, core.NewQModelIndex()), m.Index(len(m.modelData)-1, 1, core.NewQModelIndex()), []int{FirstName, LastName})
}

func (m *CustomTableModel) openfile(filePath string) {
	SetModalInstance(m)
	OpenFile(filePath)
}
