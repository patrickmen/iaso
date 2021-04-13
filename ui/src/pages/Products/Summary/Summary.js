import React, { Component } from 'react';
import 'github-markdown-css';
import ReactMarkdown from 'react-markdown';
import CssBaseline from '@material-ui/core/CssBaseline';
import Grid from '@material-ui/core/Grid';
import Container from '@material-ui/core/Container';
import HeadFeaturedPost from '@/components/Article/HeadFeaturedPost';
import Exception404 from '@/pages/ExceptionBeta/E404';
import { data } from './data';

export default class Summary extends Component {
  state = {
    markdown: [],
  }
  componentDidMount() {
    data.map((item) => {
      fetch(item)
        .then(res => res.text())
        .then(text => this.setState({markdown: [...this.state.markdown, text]}));
    })    
  }

  render() {
    const { markdown } = this.state;
    const headFeaturedPost = {
      title: 'MEET LOFLY BIO',
      description:
        "A Biopharmaceutical company, devoted to help the general public and investors better.",
      image: 'https://source.unsplash.com/random',
      imgText: 'head image description',
    };
 
    return (
      <React.Fragment>
        <CssBaseline />
        <div>
          <HeadFeaturedPost post={headFeaturedPost} />
        </div>
        <Container maxWidth="lg">
          <main>
          { markdown.length ? 
            <Grid container>
              { markdown.map((post) => (
                <ReactMarkdown
                  className="markdown-body"
                  source={post}
                  key={post.substring(0, 40)}
                  escapeHtml={false}
                />
              ))}
            </Grid> : <Exception404/> 
          }
          </main>
        </Container>
      </React.Fragment>
    );
  }
}

