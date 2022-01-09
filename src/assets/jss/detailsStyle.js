import {hexToRgb} from "@material-ui/core";

const grayColor = [
    "#999",
    "#777",
    "#3C4858",
    "#AAAAAA",
    "#D2D2D2",
    "#DDD",
    "#b4b4b4",
    "#555555",
    "#333",
    "#a9afbb",
    "#eee",
    "#e7e7e7",
];

const tooltipStyle = {
    tooltip: {
        padding: "10px 15px",
        minWidth: "130px",
        lineHeight: "1.7em",
        border: "none",
        borderRadius: "3px",
        boxShadow:
            "0 8px 10px 1px rgba(" +
            hexToRgb("#000") +
            ", 0.14), 0 3px 14px 2px rgba(" +
            hexToRgb("#000") +
            ", 0.12), 0 5px 5px -3px rgba(" +
            hexToRgb("#000") +
            ", 0.2)",
        maxWidth: "200px",
        textAlign: "center",
        fontFamily: '"Helvetica Neue",Helvetica,Arial,sans-serif',
        fontSize: "12px",
        fontStyle: "normal",
        fontWeight: "400",
        textShadow: "none",
        textTransform: "none",
        letterSpacing: "normal",
        wordBreak: "normal",
        wordSpacing: "normal",
        wordWrap: "normal",
        whiteSpace: "normal",
        lineBreak: "auto",
    },
};

const defaultFont = {
    fontFamily: '"Roboto", "Helvetica", "Arial", sans-serif',
    fontWeight: "300",
    lineHeight: "1.5em",
};

const primaryColor = "#9c27b0";
const dangerColor = "#f44336";
const detailsStyle = {
    ...tooltipStyle,
    table: {
        marginBottom: "0",
        overflow: "visible",
    },
    tableRow: {
        position: "relative",
        borderBottom: "1px solid " + grayColor[5],
    },
    tableActions: {
        display: "flex",
        border: "none",
        padding: "12px 8px !important",
        verticalAlign: "middle",
    },
    tableCell: {
        ...defaultFont,
        padding: "8px",
        verticalAlign: "middle",
        border: "none",
        lineHeight: "1.42857143",
        fontSize: "14px",
    },
    tableCellRTL: {
        textAlign: "right",
    },
    tableActionButton: {
        width: "27px",
        height: "27px",
        padding: "0",
    },
    tableActionButtonIcon: {
        width: "17px",
        height: "17px",
    },
    edit: {
        backgroundColor: "transparent",
        color: primaryColor[0],
        boxShadow: "none",
    },
    close: {
        backgroundColor: "transparent",
        color: dangerColor[0],
        boxShadow: "none",
    },
};
export default detailsStyle;
