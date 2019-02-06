import React, {Component} from "react";
import axios from "axios";
import ColorCard from "./ColorCard";
import AddColorCard from "./AddColorCard";

class SideBar extends Component {
    constructor(props) {
        super(props);
        this.state = {
            colorList: []
        };
        this.updateColorList = this.updateColorList.bind(this);
    }

    componentDidMount() {
        // TODO: remove domain when releasing.
        axios.get("http://localhost:5000/api/v1/colors/keys").then(this.updateColorList);
    }

    updateColorList({data}) {
        this.setState({colorList: data});
    }

    render() {
        console.log("rendering sidebar");
        // FIXME: make the search box work.
        let colorList = [];
        let langSet = new Set();
        for (let v of this.state.colorList) {
            colorList.push(<ColorCard lang={v.lang} name={v.name} code={v.base_code} key={v.lang + ":" + v.name}/>);
            langSet.add(v.lang);
        }
        let langList = [];
        for (let v of langSet) {
            langList.push(<div className="dropdown-item" key={v}>{v}</div>);
        }

        return (
            <div className={this.props.className}>
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

export default SideBar;
