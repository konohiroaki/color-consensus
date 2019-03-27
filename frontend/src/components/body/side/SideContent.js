import React, {Component} from "react";
import AddColorCard from "./AddColorCard";
import {isSameColor, isUndefinedColor} from "../../common/Utility";
import {actions as colors} from "../../../modules/colors/colors";
import {connect} from "react-redux";

class SideContent extends Component {

    constructor(props) {
        super(props);
        this.state = {
            langFilter: "",
            nameFilter: "",
        };

        this.SearchBar = this.SearchBar.bind(this);
        this.LangFilterSelector = this.LangFilterSelector.bind(this);
        this.NameFilterInput = this.NameFilterInput.bind(this);
        this.Cards = this.Cards.bind(this);
        this.TargetColorCard = this.TargetColorCard.bind(this);
        this.SelectableColorCards = this.SelectableColorCards.bind(this);
        this.ColorCard = this.ColorCard.bind(this);
    }

    componentDidMount() {
        this.props.fetchColors();
    }

    componentDidUpdate() {
        if (this.props.displayedColor === null) {
            this.props.setDisplayedColor();
        }
    }

    render() {
        console.log("rendering side content", this.props.colorList.length, this.props.displayedColor);
        return <div className={this.props.className} style={this.props.style}>
            <this.SearchBar/>
            <this.Cards/>
        </div>;
    }

    SearchBar() {
        return <div className="input-group rounded-0">
            <this.LangFilterSelector/>
            <this.NameFilterInput/>
        </div>;
    }

    LangFilterSelector() {
        const langOptions = getLangList(this.props.colorList)
            .map(lang => <option key={lang} value={lang}>{lang !== "" ? lang : "Language"}</option>);

        return <div className="input-group-prepend">
            <select className="custom-select rounded-0" value={this.state.langFilter}
                    onChange={e => this.setState({langFilter: e.target.value})}>
                {langOptions}
            </select>
        </div>;
    }

    NameFilterInput() {
        return <input type="text" className="form-control" value={this.state.nameFilter}
                      onChange={e => this.setState({nameFilter: e.target.value})}
        />;
    }

    Cards() {
        return <div style={{overflowY: "auto", height: "100%"}}>
            <this.TargetColorCard/>
            <this.SelectableColorCards/>
            <AddColorCard/>
        </div>;
    }

    TargetColorCard() {
        if (this.props.displayedColor === null) {
            return null;
        }
        return <div className="d-block m-2 card btn bg-dark text-light border border border-primary">
            <div className="row">
                <div className="col-3 border-right border-secondary p-3">{this.props.displayedColor.lang}</div>
                <div className="col-9 p-3">{this.props.displayedColor.name}</div>
            </div>
        </div>;
    }

    SelectableColorCards() {
        const selectableCards = this.props.colorList
            .filter(c => this.props.displayedColor !== null && !isSameColor(c, this.props.displayedColor))
            .filter(c => isLangMatchingFilter(c.lang, this.state.langFilter))
            .filter(c => isNameMatchingFilter(c.name, this.state.nameFilter))
            .sort(colorComparator)
            .map(c => <this.ColorCard key={c.lang + ":" + c.name} color={c}/>);

        return <div>{selectableCards}</div>;
    }

    ColorCard({color}) {
        return <div className="d-block m-2 card btn bg-dark text-light border border border-secondary"
                    onClick={() => this.props.setDisplayedColor(color)}>
            <div className="row">
                <div className="col-3 border-right border-secondary p-3">{color.lang}</div>
                <div className="col-9 p-3">{color.name}</div>
            </div>
        </div>;
    }
}

const getLangList = colorList => colorList.map(color => color.lang)
    .reduce((acc, current) => {
        if (!acc.includes(current)) {
            acc.push(current);
        }
        return acc;
    }, [""]);

const isLangMatchingFilter = (lang, filter) => filter === "" || lang === filter;
const isNameMatchingFilter = (name, filter) => filter === "" || name.includes(filter.toLowerCase());
const colorComparator = (c1, c2) => c1.lang !== c2.lang ? (c1.lang > c2.lang ? 1 : -1) : (c1.name > c2.name ? 1 : -1);

const mapStateToProps = state => ({
    colorList: state.colors.colors,
    displayedColor: state.colors.displayedColor,
});

const mapDispatchToProps = dispatch => ({
    fetchColors: () => dispatch(colors.fetchColors()),
    setDisplayedColor: color => dispatch(colors.setDisplayedColor(color)),
});

export default connect(mapStateToProps, mapDispatchToProps)(SideContent);
