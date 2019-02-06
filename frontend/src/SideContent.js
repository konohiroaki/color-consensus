import React, {Component} from "react";
import axios from "axios";
import ColorCard from "./ColorCard";
import AddColorCard from "./AddColorCard";

class SideContent extends Component {

    // TODO: allow update when adding color from AddColorCard? or can update locally.
    shouldComponentUpdate() {
        return this.state.colorList.length === 0;
    }

    constructor(props) {
        super(props);
        this.state = {
            colorList: []
        };
    }

    componentDidMount() {
        this.updateColorList();
    }

    updateColorList() {
        // TODO: remove domain when releasing.
        axios.get("http://localhost:5000/api/v1/colors/keys").then(({data}) => {
            this.setState({colorList: data});
            console.log(this.state);
            this.props.setTarget(this.state.colorList[1]);
        });
    }

    render() {
        console.log("rendering sidebar");
        let colorList = [];
        let langSet = new Set();
        for (let v of this.state.colorList) {
            colorList.push(
                <ColorCard lang={v.lang} name={v.name} code={v.code} setTarget={this.props.setTarget} key={v.lang + ":" + v.name}/>
            );
            langSet.add(v.lang);
        }
        let langList = [];
        for (let v of langSet) {
            langList.push(
                <div className="dropdown-item" key={v}>{v}</div>
            );
        }

        return (
            <div className={this.props.className} style={this.props.style}>
                {/* FIXME: make the search box work. */}
                <div className="input-group">
                    <button className="btn btn-outline-secondary dropdown-toggle" type="button" data-toggle="dropdown">Language</button>
                    <div className="dropdown-menu">
                        {langList}
                    </div>
                    <input type="text" className="form-control"/>
                </div>
                <div style={{overflowY: "auto", height: "100%"}}>
                    {colorList}
                    <AddColorCard/>
                </div>
            </div>
        );
    }
}

export default SideContent;
