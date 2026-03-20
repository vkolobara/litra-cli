import QtQuick
import Quickshell
import qs.Common
import qs.Services
import qs.Widgets
import qs.Modules.Plugins

PluginComponent {
    id: root

    property bool isOn: false
    property int brightness: 0
    property int temperature: 3200

    function runLitra(args) {
        Proc.runCommand("litra " + args[0], ["litra"].concat(args), null)
    }

    function toggle() {
        isOn = !isOn
        runLitra(isOn ? ["on"] : ["off"])
    }

    function setBrightness(value) {
        brightness = value
        runLitra(["brightness", String(value)])
    }

    function setTemperature(value) {
        temperature = value
        runLitra(["temperature", String(value)])
    }

    horizontalBarPill: Component {
        Row {
            spacing: Theme.spacingXS

            DankIcon {
                name: "light_mode"
                size: Theme.iconSize
                filled: root.isOn
                color: root.isOn ? Theme.primary : Theme.surfaceVariantText
                anchors.verticalCenter: parent.verticalCenter
            }

            StyledText {
                text: root.isOn ? root.brightness + "%" : I18n.tr("Off", "Litra light off status shown in the bar")
                color: root.isOn ? Theme.surfaceText : Theme.surfaceVariantText
                anchors.verticalCenter: parent.verticalCenter
            }
        }
    }

    popoutContent: Component {
        PopoutComponent {
            id: popout

            headerText: I18n.tr("Litra", "Litra light control popout header")
            showCloseButton: true

            Column {
                width: parent.width
                spacing: 0

                Rectangle {
                    width: parent.width
                    height: 1
                    color: Theme.outlineStrong
                }

                Item {
                    width: parent.width
                    height: 52

                    Row {
                        anchors.left: parent.left
                        anchors.leftMargin: Theme.spacingM
                        anchors.verticalCenter: parent.verticalCenter
                        spacing: Theme.spacingM

                        DankIcon {
                            name: "light_mode"
                            size: Theme.iconSize
                            filled: root.isOn
                            color: root.isOn ? Theme.primary : Theme.surfaceVariantText
                            anchors.verticalCenter: parent.verticalCenter
                        }

                        StyledText {
                            text: I18n.tr("Light", "Litra light power toggle label")
                            font.pixelSize: Theme.fontSizeMedium
                            anchors.verticalCenter: parent.verticalCenter
                        }
                    }

                    DankToggle {
                        anchors.right: parent.right
                        anchors.rightMargin: Theme.spacingM
                        anchors.verticalCenter: parent.verticalCenter
                        checked: root.isOn
                        onClicked: root.toggle()
                    }
                }

                Rectangle {
                    width: parent.width
                    height: 1
                    color: Theme.outlineStrong
                }

                Column {
                    width: parent.width
                    padding: Theme.spacingM
                    spacing: Theme.spacingXS
                    opacity: root.isOn ? 1.0 : 0.4

                    Behavior on opacity {
                        NumberAnimation {
                            duration: Theme.shortDuration
                            easing.type: Theme.standardEasing
                        }
                    }

                    Row {
                        spacing: Theme.spacingXS

                        DankIcon {
                            name: "brightness_high"
                            size: Theme.iconSizeSmall
                            color: Theme.surfaceVariantText
                            anchors.verticalCenter: parent.verticalCenter
                        }

                        StyledText {
                            text: I18n.tr("Brightness", "Litra brightness slider label")
                            font.pixelSize: Theme.fontSizeSmall
                            color: Theme.surfaceVariantText
                            anchors.verticalCenter: parent.verticalCenter
                        }
                    }

                    DankSlider {
                        width: parent.width - parent.padding * 2
                        minimum: 0
                        maximum: 100
                        step: 1
                        value: root.brightness
                        unit: "%"
                        enabled: root.isOn
                        leftIcon: "brightness_low"
                        rightIcon: "brightness_high"

                        onSliderDragFinished: (val) => root.setBrightness(val)
                    }
                }

                Rectangle {
                    width: parent.width
                    height: 1
                    color: Theme.outlineStrong
                }

                Column {
                    width: parent.width
                    padding: Theme.spacingM
                    spacing: Theme.spacingXS
                    opacity: root.isOn ? 1.0 : 0.4

                    Behavior on opacity {
                        NumberAnimation {
                            duration: Theme.shortDuration
                            easing.type: Theme.standardEasing
                        }
                    }

                    Row {
                        spacing: Theme.spacingXS

                        DankIcon {
                            name: "thermometer"
                            size: Theme.iconSizeSmall
                            color: Theme.surfaceVariantText
                            anchors.verticalCenter: parent.verticalCenter
                        }

                        StyledText {
                            text: I18n.tr("Temperature", "Litra color temperature slider label")
                            font.pixelSize: Theme.fontSizeSmall
                            color: Theme.surfaceVariantText
                            anchors.verticalCenter: parent.verticalCenter
                        }
                    }

                    DankSlider {
                        width: parent.width - parent.padding * 2
                        minimum: 2700
                        maximum: 6500
                        step: 100
                        value: root.temperature
                        unit: "K"
                        enabled: root.isOn
                        leftIcon: "wb_incandescent"
                        rightIcon: "wb_sunny"

                        onSliderDragFinished: (val) => root.setTemperature(val)
                    }
                }
            }
        }
    }

    popoutWidth: 400
    popoutHeight: 280
}
