var React = require('react');

var MasterReleases = React.createClass({
    render() {
        var listNode = null;
        if (this.props.releases.length > 0) {
            var releaseNodes = this.props.releases.map(release => {
                return (
                    <tr>
                        <td>{release.year}</td>
                        <td>{release.country}</td>
                    </tr>
                );
            });

            listNode = (
                <table>
                    <tbody>
                        {releaseNodes}
                    </tbody>
                </table>
            );
        } else {
            listNode = <p>No release found.</p>
        }

        return (
            <div className="master__releases">
                <h4 className="master__releases__title">Releases</h4>
                {listNode}
            </div>
        );
    }
});

module.exports = MasterReleases;