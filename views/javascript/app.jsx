class App extends React.Component {
    render() {
        return (<Home />);
    }
}

class Home extends React.Component {
    render() {
        return (
            <div className="container">
                <div className="col-xs-8 col-xs-offset-2 jumbotron text-center">
                    <h1>Home</h1>
                    <p>This is the home page</p>
                </div>
            </div>
        );
    }
}

ReactDOM.render(<App />, document.getElementById('app'));