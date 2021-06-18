import React from 'react';
import Plot from 'react-plotly.js';
import axios from 'axios';

axios.defaults.headers.post['Access-Control-Allow-Origin'] = '*';

export default class Example extends React.Component {
        state = {
            tempLine: {
                x: [],
                y: [],
                name: 'Temperature'
            },
            humidityLine: {
                x: [],
                y: [],
                name: 'Humidity'
            },
            pressureLine: {
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
        cleanState = s => {
          s.tempLine.x = [];
          s.tempLine.y = [];
          s.humidityLine.x = [];
          s.humidityLine.y = [];
          s.pressureLine.x = [];
          s.pressureLine.y = [];
        }
        rand = () => parseInt(Math.random() * 10 + this.state.revision, 10);
        increaseGraphic = () => {
            axios.get('http://127.0.0.1:8080/')
                .then(response => {
                    var all_sensor_data = response.data.data;

                    this.cleanState(this.state);
                    const { tempLine, humidityLine, pressureLine, layout } = this.state;

                    all_sensor_data.forEach(x => {
                      var dateTime = new Date(parseInt(x.Time*1000));
                      tempLine.x.push(dateTime);
                      tempLine.y.push(x.Temperature)
                      humidityLine.x.push(dateTime);
                      humidityLine.y.push(x.Humidity);
                      pressureLine.x.push(dateTime);
                      pressureLine.y.push(x.Pressure);
                    });

                    console.log(this.state);
                    // line1.x.push(all_sensor_data['Time']);
                    // line1.y.push(all_sensor_data['Temperature']);
                    // if (line1.x.length >= 100) {
                    //     line1.x.shift();
                    //     line1.y.shift();
                    // }
                    // line2.x.push(all_sensor_data['Time']);
                    // line2.y.push(all_sensor_data['Humidity']);
                    // if (line2.x.length >= 100) {
                    //     line2.x.shift();
                    //     line2.y.shift();
                    // }
                    // line3.x.push(all_sensor_data['Time']);
                    // line3.y.push(all_sensor_data['Pressure'] / 100);
                    // if (line3.x.length >= 100) {
                    //     line3.x.shift();
                    //     line3.y.shift();
                    // }
                    this.setState({ revision: this.state.revision + 1 });
                    layout.datarevision = this.state.revision + 1;
                });
        }
        render() {
            return ( <div>
                <
                Plot data = {
                    [
                        this.state.tempLine,
                        this.state.humidityLine,
                        this.state.pressureLine,
                    ]
                }
                layout = { this.state.layout }
                revision = { this.state.revision }
                graphDiv = "graph" /
                >
                </div>);
            }
        }