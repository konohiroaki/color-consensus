import React, {Component} from "react";

class App extends Component {

    render() {
        return (
            <div className="bg-dark text-light" style={{display: "flex", flexDirection: "column", height: "100%"}}>
                <Header style={{flex: "0 0 80px"}}/>
                <div style={{flex: "1 1 auto", display: "flex", flexDirection: "row"}}>
                    <div>
                        asdfsdfsadfasdfsdfsadfasdfsdfsadfasdfsdfsadf
                        asdfsdfsadfasdfsdfsadfasdfsdfsadf
                    </div>
                    <div style={{borderLeft: "#999 solid 1px", flex: "0 0 20em"}}>
                        asdfaskdjfa;slkdjf;alskdjf;alksdjf;alksdjf;alskdjf;kj
                    </div>
                </div>
            </div>
        );
    }
}

class Header extends Component {
    render() {
        return (
            <nav className="navbar navbar-dark bg-dark border-bottom border-primary" style={this.props.style}>
                <a className="navbar-brand" href="/">Color Consensus</a>
            </nav>
        );
    }
}

export default App;
