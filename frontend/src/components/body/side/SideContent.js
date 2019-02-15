import React, {Component} from "react";
import axios from "axios";
import ColorCard from "./ColorCard";
import AddColorCard from "./AddColorCard";

class SideContent extends Component {

    constructor(props) {
        super(props);
        this.state = {
            colorList: [],
            nameFilter: "",
        };
        this.langList = [];

        this.updateColorList = this.updateColorList.bind(this);
    }

    componentDidMount() {
        this.updateColorList();
    }

    updateColorList() {
        // TODO: remove domain when releasing.
        axios.get("http://localhost:5000/api/v1/colors/keys").then(({data}) => {
            console.log("side content got color list from server: ", data);
            this.langList = SideContent.getLangList(data);
            this.setState({colorList: data});

            // TODO: select random color in user's language?
            this.props.setTarget(this.state.colorList[0]);
        });
    }

    static getLangList(data) {
        return data.map(color => color.lang)
            .reduce((acc, current) => {
                if (!acc.includes(current)) {
                    acc.push(current);
                }
                return acc;
            }, []);
    }

    render() {
        console.log("rendering side content");
        const colorList = this.state.colorList
            .filter(color => this.state.nameFilter === "" || color.name.includes(this.state.nameFilter.toLowerCase()))
            .map(color => <ColorCard key={color.lang + ":" + color.name} color={color}
                                     style={{display: "block"}} setTarget={this.props.setTarget}/>);
        const langList = this.langList.map(lang => <div className="dropdown-item" key={lang}>{lang}</div>);

        return (
            <div className={this.props.className} style={this.props.style}>
                {/* TODO: make the lang filter work */}
                <div className="input-group">
                    <button className="btn btn-outline-secondary dropdown-toggle" type="button" data-toggle="dropdown">Language</button>
                    <div className="dropdown-menu">
                        {langList}
                    </div>
                    <input type="text" className="form-control" value={this.state.nameFilter}
                           onChange={e => this.setState({nameFilter: e.target.value})}/>
                </div>
                <div style={{overflowY: "auto", height: "100%"}}>
                    {colorList}
                    <AddColorCard updateColorList={this.updateColorList}/>
                </div>
            </div>
        );
    }
}

export default SideContent;
