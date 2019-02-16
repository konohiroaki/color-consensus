import React, {Component} from "react";
import axios from "axios";

class Header extends Component {

    constructor(props) {
        super(props);
        this.state = {
            userId: null,
        };
    }

    render() {
        const userId = this.state.userId;
        const button = userId === null || userId === ""
                       ? null
            // TODO: click button to copy userId into clipboard
                       : <button className="btn btn-outline-secondary">ID: {userId}</button>;
        console.log("rendering header");
        return (
            <nav className="navbar navbar-dark bg-dark border-bottom border-secondary" style={this.props.style}>
                <a className="navbar-brand" href="/">Color Consensus</a>
                {button}
            </nav>
        );
    }

    componentDidMount() {
        axios.get("http://localhost:5000/api/v1/users/presence")
            .then(({data}) => {
                // TODO: show user id on header
                console.log("user present", data);
                this.setState({userId: data.userID});
            })
            .catch(() => {
                // TODO: show modal to login or sign up.
                console.log("user not present");
                this.setState({userId: ""});
            });
    }
}

export default Header;