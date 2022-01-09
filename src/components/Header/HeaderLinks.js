import React from "react";
import {makeStyles} from "@material-ui/core/styles";
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import Tooltip from "@material-ui/core/Tooltip";
import {CloudDownload} from "@material-ui/icons";
import styles from "../../assets/jss/headerLinksStyle.js";
import Button from "@material-ui/core/Button";
import YouTubeIcon from '@mui/icons-material/YouTube';
import LinkedInIcon from '@mui/icons-material/LinkedIn';
import GitHubIcon from '@mui/icons-material/GitHub';

const useStyles = makeStyles(styles);

export default function HeaderLinks(props) {
    const classes = useStyles();
    return (
        <List className={classes.list}>
            <ListItem className={classes.listItem}>
                <Button
                    href="https://drive.google.com/file/d/1uzAXRxNnkE677j8o6qoKBgH6cqWKlLaC/view?usp=sharing"
                    color="transparent"
                    target="_blank"
                    className={classes.navLink}
                >
                    <CloudDownload className={classes.icons}/> Resume Preview
                </Button>
            </ListItem>
            <ListItem className={classes.listItem}>
                <Tooltip
                    id="instagram-facebook"
                    title="My LinkedIn"
                    placement={window.innerWidth > 959 ? "top" : "left"}
                    classes={{tooltip: classes.tooltip}}
                >
                    <Button
                        color="transparent"
                        href="https://www.linkedin.com/in/yildirimtyln/"
                        target="_blank"
                        className={classes.navLink}
                        startIcon={<LinkedInIcon/>}

                    >
                        <i className={classes.socialIcons + " fab fa-facebook"}/>
                    </Button>
                </Tooltip>
            </ListItem>
            <ListItem className={classes.listItem}>
                <Tooltip
                    id="instagram-tooltip"
                    title="My GitHub"
                    placement={window.innerWidth > 959 ? "top" : "left"}
                    classes={{tooltip: classes.tooltip}}
                >
                    <Button
                        color="transparent"
                        href="https://github.com/TaylanYildirim"
                        target="_blank"
                        className={classes.navLink}
                        startIcon={<GitHubIcon/>}
                    >
                        <i className={classes.socialIcons + " fab fa-instagram"}/>
                    </Button>
                </Tooltip>
            </ListItem>
            <ListItem className={classes.listItem}>
                <Tooltip
                    id="instagram-twitter"
                    title="My YouTube"
                    placement={window.innerWidth > 959 ? "top" : "left"}
                    classes={{tooltip: classes.tooltip}}
                >
                    <Button
                        href="https://www.youtube.com/watch?v=vzoM75_d61Q&ab_channel=TaylanY%C4%B1ld%C4%B1r%C4%B1m"
                        target="_blank"
                        color="transparent"
                        className={classes.navLink}
                        startIcon={<YouTubeIcon/>}
                    >
                        <i className={classes.socialIcons + " fab fa-twitter"}/>
                    </Button>
                </Tooltip>
            </ListItem>
        </List>
    );
}
