var React = require('react');


var Pager = React.createClass({
    _onPageChange(page) {
        this.props.handler(page);
    },

    render() {
        var prevNode = null;
        var nextNode = null;

        console.log(this.props.pager);

        if (this.props.pager.page > 1) {
            prevNode = (
                <span className="pager__prev" onClick={() => (this._onPageChange(this.props.pager.page - 1))}>
                    <span className="pager__icon"><i className="fa fa-angle-left"/></span>
                    <span className="pager__label pager__label--prev">previous</span>
                </span>
            );
        }

        if (this.props.pager.page < this.props.pager.page_count) {
            nextNode = (
                <span className="pager__next" onClick={() => (this._onPageChange(this.props.pager.page + 1))}>
                    <span className="pager__label pager__label--next">next</span>
                    <span className="pager__icon"><i className="fa fa-angle-right"/></span>
                </span>
            );
        }

        return (
            <div className="pager">
                {prevNode}
                <span className="pager__page-progress">
                    <span className="pager__page-progress__current">{this.props.pager.page}</span>
                    <span className="pager__page-progress__sep">/</span>
                    <span className="pager__page-progress__total">{this.props.pager.page_count}</span>
                </span>
                {nextNode}
            </div>
        );
    }
});

module.exports = Pager;