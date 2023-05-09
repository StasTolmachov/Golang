@extends('layout')

@section('title')
About
@endsection

@section('main_content')
<h1>Review</h1>

<form method="post" action="/review/check">
@csrf
    <input type="email" name="email" id="email" class="form-control"><br>
    <input type="text" name="subject" id="subject" class="form-control"><br>
    <textarea name="message" id="message" class="form-control"></textarea><br>
    <button type="submit" class="btn btn-success">Send</button>
</form>

@endsection
