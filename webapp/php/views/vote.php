<div class="jumbotron">
    <div class="container">
        <h1>清き一票をお願いします！！！</h1>
    </div>
</div>
<div class="container">
    <div class="row">
        <div class="col-md-6 col-md-offset-3">
            <div class="login-panel panel panel-default">
                <div class="panel-heading">
                    <h3 class="panel-title">投票フォーム</h3>
                </div>
                <div class="panel-body">
                    <form method="POST" action="/vote">
                        <fieldset>
                            <label>氏名</label>
                            <div class="form-group">
                                <input class="form-control" name="name" autofocus>
                            </div>
                            <label>住所</label>
                            <div class="form-group">
                                <input class="form-control" name="address" value="">
                            </div>
                            <label>私の番号</label>
                            <div class="form-group">
                                <input class="form-control" name="mynumber" value="">
                            </div>
                            <label>候補者</label>
                            <div class="form-group">
                                <select name="candidate">
                                    <?php foreach ($candidates as $candidate) { ?>
                                        <option value="<?= $candidate['name'] ?>"><?= $candidate['name'] ?></option>
                                    <?php } ?>
                                </select>
                            </div>
                            <label>投票理由</label>
                            <div class="form-group">
                                <input class="form-control" name="keyword" value="">
                            </div>
                            <label>投票数</label>
                            <div class="form-group">
                                <input class="form-control" name="vote_count" value="">
                            </div>

                            <div class="text-danger"><?= $message ?></div>
                            <input class="btn btn-lg btn-success btn-block" type="submit" name="vote" value="投票"/>
                        </fieldset>
                    </form>
                </div>
            </div>
        </div>
    </div>
</div>
