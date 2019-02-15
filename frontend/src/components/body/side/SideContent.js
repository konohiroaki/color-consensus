import React, {Component} from "react";
import axios from "axios";
import ColorCard from "./ColorCard";
import AddColorCard from "./AddColorCard";

class SideContent extends Component {

    constructor(props) {
        super(props);
        this.state = {
            colorList: [],
            langFilter: "",
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
            }, [""]);
    }

    render() {
        console.log("rendering side content");
        const langList = this.langList.map(lang => <option key={lang} value={lang}>{lang !== "" ? lang : "Language"}</option>);
        const colorList = this.state.colorList
            .filter(color => this.state.nameFilter === "" || color.name.includes(this.state.nameFilter.toLowerCase()))
            .filter(color => this.state.langFilter === "" || color.lang === this.state.langFilter)
            .sort((a, b) => SideContent.colorComparator(a, b))
            .map(color => <ColorCard key={color.lang + ":" + color.name} color={color}
                                     style={{display: "block"}} setTarget={this.props.setTarget}/>);

        return (
            <div className={this.props.className} style={this.props.style}>
                <div className="input-group" style={{borderRadius: "0px"}}>
                    <div className="input-group-prepend">
                        <select className="custom-select" value={this.state.langFilter} style={{borderRadius: "0"}}
                                onChange={e => this.setState({langFilter: e.target.value})}>
                            {langList}
                        </select>
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

    // ascending order for lang -> name
    static colorComparator(a, b) {
        if (a.lang === b.lang) {

            return a.name < b.name ? -1 : 1;
        }
        return a.lang < b.lang ? -1 : 1;
    }
}

export default SideContent;
