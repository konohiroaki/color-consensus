import React, {Component} from "react";
import axios from "axios";
import ColorCard from "./ColorCard";
import AddColorCard from "./AddColorCard";

class SideContent extends Component {

    constructor(props) {
        super(props);
        this.state = {
            colorList: [],
            searchText: "",
        };
        this.handleSearchChange = this.handleSearchChange.bind(this);
    }

    componentDidMount() {
        this.updateColorList();
    }

    updateColorList() {
        // TODO: remove domain when releasing.
        axios.get("http://localhost:5000/api/v1/colors/keys").then(({data}) => {
            console.log("side content got color list from server: ", data);
            this.setState({colorList: data});

            // TODO: select random color in user's language?
            this.props.setTarget(this.state.colorList[0]);
        });
    }

    handleSearchChange(event) {
        let value = event.target.value;
        this.setState({searchText: value});
    }

    render() {
        console.log("rendering side content");
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
                {/* https://www.w3schools.com/jquery/tryit.asp?filename=tryjquery_filters_anything */}
                <div className="input-group">
                    <button className="btn btn-outline-secondary dropdown-toggle" type="button" data-toggle="dropdown">Language</button>
                    <div className="dropdown-menu">
                        {langList}
                    </div>
                    <input type="text" className="form-control" value={this.state.searchText} onChange={this.handleSearchChange}/>
                </div>
                <div style={{overflowY: "auto", height: "100%"}}>
                    <div id="colorList">
                        {colorList}
                    </div>
                    <AddColorCard/>
                </div>
            </div>
        );
    }
}

export default SideContent;
