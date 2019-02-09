import React, {Component} from "react";

class Header extends Component {

    shouldComponentUpdate() {
        return false;
    }

    render() {
        return (
            <nav className="navbar navbar-dark bg-dark border-bottom border-secondary" style={this.props.style}>
                <a className="navbar-brand" href="/">Color Consensus</a>
            </nav>
        );
    }
}

export default Header;