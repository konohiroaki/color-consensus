import React, {Component} from "react";
import {Route, Switch} from "react-router-dom";
import VotingPage from "./vote/VotingPage";
import StatisticsPage from "./statistics/StatisticsPage";

class MainContent extends Component {

    render() {
        console.log("rendering [main content]");

        return <div className="container-fluid pt-3"
                    style={Object.assign({overflowY: "auto", height: "100%"}, this.props.style)}>
            <Switch>
                <Route exact path={"/"} render={({history}) => (
                    <VotingPage history={history} loginModalRef={this.props.loginModalRef}/>
                )}/>
                <Route path={"/statistics"} render={({history}) => (
                    <StatisticsPage history={history}/>
                )}/>
            </Switch>
        </div>;
    }
}

export default MainContent;
