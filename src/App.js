import ApsisTable from "./components/Table/ApsisTable";
import {useEffect} from "react";
import Header from "./components/Header/Header";
import Parallax from "./components/Parallax/Parallax";
import Footer from "./components/Footer/Footer";
import * as PropTypes from "prop-types";
import {makeStyles} from "@material-ui/core/styles";
import HeaderLinks from "./components/Header/HeaderLinks";
import styles from "./assets/jss/headerStyle.js";
import TableDetail from "./components/Table/TableDetail";


function GridItem(props) {
    return null;
}

GridItem.propTypes = {children: PropTypes.node};

function GridContainer(props) {
    return null;
}

GridContainer.propTypes = {children: PropTypes.node};

const useStyles = makeStyles(styles);

const App = () => {
    const classes = useStyles();
    useEffect(() => {

    }, [])
    return (
        <>
            <Header
                brand="CASE STUDY"
                rightLinks={<HeaderLinks/>}
                fixed
                color="transparent"
                changeColorOnScroll={{
                    height: 400,
                    color: "white",
                }}
            />
            <Parallax image={require("./assets/image/apsis.png").default}>
                <div className={classes.container}>
                    <GridContainer>
                        <GridItem>
                            <div className={classes.brand}>
                                <h1 className={classes.title}>CASE STUDY</h1>
                                <h3 className={classes.subtitle}>
                                    A Badass Material-UI Kit based on Material Design.
                                </h3>
                            </div>
                        </GridItem>
                    </GridContainer>
                </div>
            </Parallax>
            <TableDetail/>

            <Footer/>
        </>
    );
}

export default App;
