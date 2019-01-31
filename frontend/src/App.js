import React, {Component} from "react";

class App extends Component {

    render() {
        return (
            <div className="bg-dark text-light" style={{display: "flex", flexDirection: "column", height: "100%"}}>
                <Header style={{flex: "0 0 80px"}}/>
                <div style={{flex: "1 1 auto", display: "flex", flexDirection: "row"}}>
                    <MainContent/>
                    <SideBar className="border-left border-secondary"/>
                </div>
            </div>
        );
    }
}

class Header extends Component {
    render() {
        return (
            <nav className="navbar navbar-dark bg-dark border-bottom border-secondary" style={this.props.style}>
                <a className="navbar-brand" href="/">Color Consensus</a>
            </nav>
        );
    }
}

class MainContent extends Component {
    constructor(props) {
        super(props);
        this.state = {
            // similar color 2d array
        };
    }

    componentDidMount() {
        // calculate similar colors
    }

    render() {
        // TODO: think algorithm to list similar colors.
        // TODO: think algorithm to place them.
        // TODO: understand react-selectable-fast and apply for them.
        return (
            <div className="container-fluid">
                asdf
            </div>
        );
    }
}

class SideBar extends Component {
    constructor(props) {
        super(props);
        this.state = {
            // api call result.
        };
    }

    componentDidMount() {
        // call api. axios.
    }

    render() {
        // TODO: create color card from api result.
        // TODO: when click, it should draw the main content for it. (maybe <button> is better than <a>)
        // TODO: add search box at the top for better UX.
        return (
            <div className={this.props.className} style={{overflowY: "auto", width: "20em"}}>
                <ColorCard/>
                <ColorCard/>
                <ColorCard/>
                <ColorCard/>
                <ColorCard/>
                <ColorCard/>
                <ColorCard/>
                <ColorCard/>
                <ColorCard/>
                <ColorCard/>
                <ColorCard/>
                <ColorCard/>
                <ColorCard/>
                <ColorCard/>
                <ColorCard/>
                <ColorCard/>
                <ColorCard/>
                <ColorCard/>
            </div>
        );
    }
}

class ColorCard extends Component {
    render() {
        return (
            <a className="card bg-dark border border-secondary m-2" href="/">
                <div className="card-body text-light text-center">
                    ja : 赤 (10 Votes)
                </div>
            </a>
        );
    }
}

export default App;
