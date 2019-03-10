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

        this.updateColorList = this.updateColorList.bind(this);
    }

    componentDidMount() {
        this.updateColorList().then(() => {
            // TODO: select random color in user's language?
            this.props.setTarget(this.state.colorList[0]);
        });
    }

    updateColorList() {
        return axios.get(`${process.env.WEBAPI_HOST}/api/v1/colors/keys`).then(({data}) => {
            console.log("side content got color list from server: ", data);
            this.setState({colorList: data});
        });
    }

    render() {
        console.log("rendering side content");

        return (
            <div className={this.props.className} style={this.props.style}>
                <SearchBar langList={this.state.colorList.map(color => color.lang)}
                           langFilter={this.state.langFilter}
                           langFilterSetter={e => this.setState({langFilter: e.target.value})}
                           nameFilter={this.state.nameFilter}
                           nameFilterSetter={e => this.setState({nameFilter: e.target.value})}/>
                <div style={{overflowY: "auto", height: "100%"}}>
                    <ColorCards colorList={this.state.colorList}
                                langFilter={this.state.langFilter}
                                nameFilter={this.state.nameFilter}
                                setTarget={this.props.setTarget}/>
                    <AddColorCard updateColorList={this.updateColorList}/>
                </div>
            </div>
        );
    }
}

class SearchBar extends Component {

    getLangList() {
        return this.props.langList.reduce((acc, current) => {
            if (!acc.includes(current)) {
                acc.push(current);
            }
            return acc;
        }, [""]);
    }

    render() {
        const langList = this.getLangList().map(lang => (
            <option key={lang} value={lang}>{lang !== "" ? lang : "Language"}</option>
        ));

        return (
            <div className="input-group" style={{borderRadius: "0px"}}>
                <div className="input-group-prepend">
                    <select className="custom-select" value={this.props.langFilter} style={{borderRadius: "0"}}
                            onChange={this.props.langFilterSetter}>
                        {langList}
                    </select>
                </div>
                <input type="text" className="form-control" value={this.props.nameFilter}
                       onChange={this.props.nameFilterSetter}/>
            </div>
        );
    }
}

class ColorCards extends Component {

    render() {
        const colorCards = this.props.colorList
            .filter(color => this.props.nameFilter === "" || color.name.includes(this.props.nameFilter.toLowerCase()))
            .filter(color => this.props.langFilter === "" || color.lang === this.props.langFilter)
            // TODO: sort on server side.
            .sort((a, b) => ColorCards.colorComparator(a, b))
            .map(color => (
                <ColorCard key={color.lang + ":" + color.name} color={color}
                           style={{display: "block"}} setTarget={this.props.setTarget}/>
            ));

        return (
            <div>
                {colorCards}
            </div>
        );
    }

    // ascending order for lang -> name
    static colorComparator(a, b) {
        if (a.lang === b.lang) {
            return a.name - b.name;
        }
        return a.lang - b.lang;
    }
}

export default SideContent;
