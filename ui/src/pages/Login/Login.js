import React from 'react';
import { message } from 'antd';
import Avatar from '@material-ui/core/Avatar';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import Checkbox from '@material-ui/core/Checkbox';
import Link from '@material-ui/core/Link';
import Grid from '@material-ui/core/Grid';
import Box from '@material-ui/core/Box';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';
import logo from '@/assets/img/logo.png';
import { submitLogin, errorAnalysis } from '@/services';
import styles from './Login.less';

function Copyright() {
  return (
    <Typography variant="body2" color="textSecondary" align="center">
      {'Copyright © '}
      <Link color="inherit" href="https://material-ui.com/">
        Patrick
      </Link>{' '}
      {new Date().getFullYear()}
      {'.'}
    </Typography>
  );
}

const useStyles = makeStyles((theme) => ({
  paper: {
    marginTop: theme.spacing(8),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  avatar: {
    margin: theme.spacing(1),
    backgroundColor: theme.palette.secondary.main,
  },
  form: {
    width: '100%', // Fix IE 11 issue.
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
    color: '#fff',
    backgroundColor: '#61bcca',
  },
}));

export default function SignIn(props) {
  const classes = useStyles();
  const [values, setValues] = React.useState({
    username: '',
    password: '',
  });

  const handleChange = name => event => {
    setValues({ ...values, [name]: event.target.value });
  };

  const handleValid = async(e) => {
    let isValid = true;
    let isLack = false;
    for (let value in values) {
      if (values[value] === '' || values[value] === null) {
        isValid = false;
        isLack = true;
      }
      if (!isValid) {
        break
      }
    }
    if (!isLack) {
      e.preventDefault()
      if (isValid) {
        const response = await submitLogin(values);
        document.getElementById("form").reset();
        if (response.data) {
          props.history.push(`/about-us`);
          message.success('Login success!');
        } else if (response.response) {
          let error = errorAnalysis(response.response);
          message.error('Username or password is incorrect!');
        }
      }
    }
    setValues({ ...values, ["username"]: '' });
    setValues({ ...values, ["password"]: '' });
  };

  return (
    <Container component="main" maxWidth="xs">
      <CssBaseline />
      <div className={classes.paper}>
        {/* <Avatar className={classes.avatar}>
          <LockOutlinedIcon />
        </Avatar> */}
        <img alt="logo" className={styles.logo} src={logo} />
        {/* <Typography component="h1" variant="h5">
          FLY BIO
        </Typography> */}
        <form className={classes.form} id="form" method="post">
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            id="username"
            label="Username"
            name="username"
            autoComplete="username"
            onChange={handleChange("username")}
            autoFocus
          />
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            name="password"
            label="Password"
            type="password"
            id="password"
            autoComplete="current-password"
            onChange={handleChange("password")}
          />
          <Button
            type="submit"
            fullWidth
            variant="contained"
            className={classes.submit}
            onClick={handleValid}
          >
            Sign In
          </Button>
        </form>
      </div>
      <Box mt={8}>
        <Copyright />
      </Box>
    </Container>
  );
}