import React, {Component} from "react";

class ColorCard extends Component {

    shouldComponentUpdate(nextProps) {
        return JSON.stringify(this.props.color) !== JSON.stringify(nextProps.color)
               || JSON.stringify(this.props.style) !== JSON.stringify(nextProps.style);
    }

    render() {
        console.log("rendering color card", this.props);
        return (
            <a className="card btn bg-dark border border-secondary m-2" style={this.props.style}
               onClick={() => this.props.setTarget(this.props.color)}>
                <div className="row">
                    <div className="col-3 border-right border-secondary p-3">{this.props.color.lang}</div>
                    <div className="col-9 p-3">{this.props.color.name}</div>
                </div>
            </a>
        );
    }
}

export default ColorCard;