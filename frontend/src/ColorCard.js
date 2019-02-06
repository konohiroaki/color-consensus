import React, {Component} from "react";

class ColorCard extends Component {

    handleClick(lang, name, code) {
        this.props.setTarget({lang: lang, name: name, code: code});
    }

    render() {
        console.log("rendering color card");
        return (
            <a className="card btn bg-dark border border-secondary m-2"
               onClick={this.handleClick.bind(this, this.props.lang, this.props.name, this.props.code)}>
                <div className="row">
                    <div className="col-3 border-right border-secondary p-3">{this.props.lang}</div>
                    <div className="col-9 p-3">{this.props.name}</div>
                </div>
            </a>
        );
    }
}

export default ColorCard;