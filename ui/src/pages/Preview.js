import React, { Component } from 'react';
import CssBaseline from '@material-ui/core/CssBaseline';
import Grid from '@material-ui/core/Grid';
import Container from '@material-ui/core/Container';
import HeadFeaturedPost from '@/components/Article/HeadFeaturedPost';
import PictureAlignRight from '@/components/Article/PictureAlignRight';
import PictureAlignLeft from '@/components/Article/PictureAlignLeft';
import PictureAlignJustify from '@/components/Article/PictureAlignJustify';

export default class Preview extends Component {
  state = {
    data: [],
    headPost: {},
  }
  componentDidMount() {
    this.setState({
      data: [...this.state.data, window.opener.data],
      headPost: window.opener.data.headPost,
    });   
  }

  render() {
    const { data, headPost } = this.state;
  
    return (
      <React.Fragment>
        <CssBaseline />
        <div>
          <HeadFeaturedPost post={headPost} />
        </div>
        <Container maxWidth="lg"> 
          <main>
            <Grid container>
              { data.map((post) => (
                <div key={JSON.parse(post.content).substring(0, 40)}>
                  <div>
                    {post.align == "right" ? <PictureAlignRight post={post} /> : post.align == "left" ? <PictureAlignLeft post={post} /> : <PictureAlignJustify post={post} />}
                  </div>
                </div>
              ))}
            </Grid>
          </main>
        </Container>
      </React.Fragment>
    );
  }
}
