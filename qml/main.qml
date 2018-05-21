import QtQuick 2.7
import QtQuick.Controls 1.4
import QtQuick.Controls 2.1
import CustomQmlTypes 1.0

Item {
    id: window
    height: 540

    Rectangle {
        id: rectangle1
        color: "#f7f8f8"
        anchors.fill: parent

        Item {
            id: item_header
            anchors.fill: parent
            Rectangle {
                id: rectangle
                height: 50
                color: "#e23536"
                anchors.left: parent.left
                anchors.leftMargin: 0
                anchors.right: parent.right
                anchors.rightMargin: 0
                anchors.top: parent.top
                anchors.topMargin: 0

                Text {
                    id: text1
                    color: "#ffffff"
                    text: qsTr("Jio Phone IMEI Lookup")
                    anchors.horizontalCenter: parent.horizontalCenter
                    anchors.bottom: parent.bottom
                    anchors.top: parent.top
                    textFormat: Text.PlainText
                    elide: Text.ElideRight
                    style: Text.Normal
                    renderType: Text.NativeRendering
                    clip: false
                    font.bold: true
                    font.family: "Arial"
                    lineHeight: 0.5
                    verticalAlignment: Text.AlignVCenter
                    horizontalAlignment: Text.AlignHCenter
                    font.pixelSize: 23
                }
            }
            Item {
                id: item_header1
                width: 640
                height: 45
                anchors.top: rectangle.bottom
                anchors.topMargin: 4

                Label {
                    id: label
                    text: qsTr("File :")
                    anchors.verticalCenter: parent.verticalCenter
                    anchors.left: parent.left
                    anchors.leftMargin: 25
                }

                Button {
                    id: button
                    width: 113
                    height: 40
                    text: qsTr("Select File")
                    anchors.right: parent.right
                    anchors.rightMargin: 22
                    anchors.verticalCenter: parent.verticalCenter
                    font.family: "Arial"
                    anchors.leftMargin: 30
                    focusPolicy: Qt.StrongFocus
                    onClicked: tableview.model.openfile("sample.xlsx")
                }

                TextField {
                    id: textField
                    anchors.verticalCenter: parent.verticalCenter
                    anchors.right: button.left
                    anchors.rightMargin: 20
                    anchors.left: label.right
                    anchors.leftMargin: 15
                    placeholderText: qsTr("Text Field")
                }
            }

            TableView {
                id: tableview
                height: 365
                anchors.top: item_header1.bottom
                anchors.topMargin: 6
                anchors.right: parent.right
                anchors.rightMargin: 8
                anchors.left: parent.left
                anchors.leftMargin: 8
                transformOrigin: Item.Center
                model: CustomTableModel{}

                TableViewColumn {
                        role: "IMEI"
                        title: role
                    }

                TableViewColumn {
                        role: "Distributor"
                        title: role
                    }
                TableViewColumn {
                        role: "Retailer"
                        title: role
                    }

                TableViewColumn {
                        role: "Head1"
                        title: role
                    }

                TableViewColumn {
                        role: "Date"
                        title: role
                    }
            }

            Item {
                id: item1
                height: 50
                anchors.top: tableview.bottom
                anchors.topMargin: 0
                anchors.left: parent.left
                anchors.leftMargin: 0
                anchors.right: parent.right
                anchors.rightMargin: 0

                Button {
                    id: button1
                    width: 150
                    text: qsTr("Start")
                    anchors.verticalCenter: parent.verticalCenter
                    spacing: 7
                    font.family: "Tahoma"
                    anchors.top: parent.top
                    anchors.topMargin: 8
                    anchors.left: parent.left
                    anchors.leftMargin: 120
                }

                Button {
                    id: button2
                    width: 150
                    text: qsTr("Update")
                    anchors.verticalCenter: parent.verticalCenter
                    font.family: "Tahoma"
                    anchors.top: parent.top
                    anchors.topMargin: 8
                    anchors.right: parent.right
                    anchors.rightMargin: 120
                }
            }

            Item {
                id: item2
                anchors.bottom: parent.bottom
                anchors.bottomMargin: 0
                anchors.top: item1.bottom
                anchors.topMargin: 0
                anchors.left: parent.left
                anchors.leftMargin: 0
                anchors.right: parent.right
                anchors.rightMargin: 0

                Label {
                    id: label1
                    text: qsTr("Created By Arpit Agarwal")
                    anchors.verticalCenter: parent.verticalCenter
                    font.capitalization: Font.MixedCase
                    font.strikeout: false
                    font.italic: false
                    font.bold: true
                    fontSizeMode: Text.HorizontalFit
                    renderType: Text.NativeRendering
                    font.pointSize: 10
                    font.family: "Tahoma"
                    textFormat: Text.AutoText
                    verticalAlignment: Text.AlignVCenter
                    horizontalAlignment: Text.AlignRight
                    anchors.right: parent.right
                    anchors.rightMargin: 10
                }
            }


        }
    }
    
}
