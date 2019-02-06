import React, {Component} from "react";

class ColorCard extends Component {
    constructor(props) {
        super(props);
        this.render.bind(this);
    }

    handleClick(lang, name, code) {
        console.log(lang, name, code);
        // TODO: somehow affect to main content.
        // draw(lang, name, code);
    }

    render() {
        console.log("rendering color card");
        return (
            <a className="card btn bg-dark border border-secondary m-2" onClick={this.handleClick.bind(this, this.props.lang, this.props.name, this.props.code)}>
                <div className="row">
                    <div className="col-3 border-right border-secondary p-3">{this.props.lang}</div>
                    <div className="col-9 p-3">{this.props.name}</div>
                </div>
            </a>
        );
    }
}

export default ColorCard;