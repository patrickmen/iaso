import React, { Component } from 'react';
import 'github-markdown-css';
import ReactMarkdown from 'react-markdown';
import CssBaseline from '@material-ui/core/CssBaseline';
import Grid from '@material-ui/core/Grid';
import Container from '@material-ui/core/Container';
import HeadFeaturedPost from '@/components/Article/HeadFeaturedPost';

export default class Preview extends Component {
  state = {
    markdown: [],
  }
  componentDidMount() {
      let content = JSON.parse(window.opener.data.content);
      this.setState({markdown: [...this.state.markdown, content]});   
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
            <Grid container>
              { markdown.map((post) => (
                <ReactMarkdown
                  className="markdown-body"
                  source={post}
                  key={post.substring(0, 40)}
                  escapeHtml={false}
                />
              ))}
            </Grid>
          </main>
        </Container>
      </React.Fragment>
    );
  }
}
