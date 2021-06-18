import React from 'react';
import Plot from 'react-plotly.js';
import axios from 'axios';

axios.defaults.headers.post['Access-Control-Allow-Origin'] = '*';

export default class Example extends React.Component {
        state = {
            line1: {
                x: [],
                y: [],
                name: 'Temperature'
            },
            line2: {
                x: [],
                y: [],
                name: 'Humidity'
            },
            line3: {
                x: [],
                y: [],
                name: 'Pressure'
            },
            layout: {
                datarevision: 0,
                width: 1200,
                height: 800,
                title: 'Israel forever'
            },
            revision: 0,
        }
        componentDidMount() {
            setInterval(this.increaseGraphic, 1000);
        }
        rand = () => parseInt(Math.random() * 10 + this.state.revision, 10);
        increaseGraphic = () => {
            axios.get('http://127.0.0.1:8080/')
                .then(response => {
                    var last = response.data.data.pop()

                    const { line1, line2, line3, layout } = this.state;
                    line1.x.push(last['Time']);
                    line1.y.push(last['Temperature']);
                    if (line1.x.length >= 100) {
                        line1.x.shift();
                        line1.y.shift();
                    }
                    line2.x.push(last['Time']);
                    line2.y.push(last['Humidity']);
                    if (line2.x.length >= 100) {
                        line2.x.shift();
                        line2.y.shift();
                    }
                    line3.x.push(last['Time']);
                    line3.y.push(last['Pressure'] / 100);
                    if (line3.x.length >= 100) {
                        line3.x.shift();
                        line3.y.shift();
                    }
                    this.setState({ revision: this.state.revision + 1 });
                    layout.datarevision = this.state.revision + 1;
                });
        }
        render() {
            return ( < div >
                <
                Plot data = {
                    [
                        this.state.line1,
                        this.state.line2,
                        this.state.line3,
                    ]
                }
                layout = { this.state.layout }
                revision = { this.state.revision }
                graphDiv = "graph" /
                >
                <
                /div>);
            }
        }