import React, {Component} from "react";
import {Route, Switch} from "react-router-dom";
import VotingPage from "./VotingPage";
import StatisticsPage from "./StatisticsPage";

class MainContent extends Component {

    render() {
        console.log("rendering main content");

        return (
            <div className="container-fluid pt-3" style={Object.assign({overflowY: "auto"}, this.props.style)}>
                <Switch>
                    <Route exact path={"/"} render={({history}) => (
                        <VotingPage userId={this.props.userId} target={this.props.target} history={history}
                                    loginModalRef={this.props.loginModalRef}/>
                    )}/>
                    <Route path={"/statistics"} render={({history}) => (
                        <StatisticsPage target={this.props.target} history={history}/>
                    )}/>
                </Switch>
            </div>
        );
    }
}

export default MainContent;