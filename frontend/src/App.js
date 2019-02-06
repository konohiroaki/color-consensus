import React, {Component} from "react";
import Header from "./Header";
import Body from "./Body";

// TODO: do test for react app (jest?)
class App extends Component {

    render() {
        return (
            <div className="bg-dark text-light" style={{display: "flex", flexDirection: "column", height: "100%"}}>
                <Header style={{flex: "0 0 80px"}}/>
                <Body style={{flex: "1 1 auto"}}/>
            </div>
        );
    }
}

export default App;
