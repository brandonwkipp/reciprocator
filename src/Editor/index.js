import React, { Component } from 'react';
import {
  Button, Col, Form, FormGroup, Label, Row,
} from 'reactstrap';

class Editor extends Component {
  constructor(props) {
    super(props);
    this.state = {
    };
  }

  render() {
    return (
      <>
        <Form>
          <Row>
            <Col md={7}>
              <FormGroup className="px-2">
                <Label for="blog-editor-title">Title</Label>
                <input id="blog-editor-title" className="w-100" type="text" name="title" value="" />
              </FormGroup>
              <FormGroup className="px-2">
                <Label for="blog-editor-body">Body</Label>
                <textarea id="blog-editor-body" className="w-100" name="body" rows="8" />
              </FormGroup>
              <FormGroup className="px-2">
                <Label for="blog-editor-tags">Preview</Label>
                <input id="blog-editor-preview" className="w-100" type="text" name="preview" value="" />
              </FormGroup>
              <FormGroup className="px-2">
                <Label for="blog-editor-tags">Tags</Label>
                <input id="blog-editor-tags" className="w-100" type="text" name="tags" value="" />
              </FormGroup>
            </Col>
            <Col md={5}>
              <FormGroup className="px-2">
                <Label for="image">Choose Image File:</Label>
                <input className="w-100" type="file" name="image" />
              </FormGroup>
              <FormGroup className="px-2">
                <Label className="switch">
                  <input id="publish-to-rt" type="checkbox" />
                  <span className="slider round" />
                </Label>
                <span className="verticalAlignTop">Publish to RadiumTree</span>
              </FormGroup>
              <FormGroup className="px-2">
                <Label className="switch">
                  <input id="publish-to-medium" type="checkbox" />
                  <span className="slider round" />
                </Label>
                <span className="verticalAlignTop">Publish to Medium</span>
              </FormGroup>
              <FormGroup className="px-2">
                <Button id="createBlogPostSubmit" className="call-to-action-button" type="button">Save post</Button>
              </FormGroup>
            </Col>
          </Row>
        </Form>
        <div id="createBlogPostToast" className="toast fade fixed-top" data-delay="5000">
          <div className="toast-header">
            <strong className="mr-auto">RadiumTree</strong>
            <button type="button" className="ml-2 mb-1 close" data-dismiss="toast" aria-label="Close">
              <span aria-hidden="true">&times;</span>
            </button>
          </div>
          <div id="createBlogPostToastBody" className="toast-body text-left" />
        </div>
      </>
    );
  }
}

export default Editor;
