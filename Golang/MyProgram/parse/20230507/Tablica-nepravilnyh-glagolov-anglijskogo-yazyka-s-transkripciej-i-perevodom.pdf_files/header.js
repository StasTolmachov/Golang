// $(function () {
//     $('.js-register-buttons > span').on('click', function(){

//         let $btn = $(this)
//         let $modal = $btn.closest('.modal-body');

//         $modal.find('.collection-button--register').hide();

//         if( $btn.hasClass('js-reg-school')){
//             $modal.find('.col-md-12').after(`
//             <div class="js-school-form">
//                 <input type="hidden" name="type" value="school">
//                 <input id="register_username" name="username" autocomplete="off" readonly onfocus="this.removeAttribute('readonly')" class="form-control form-placeholder" type="text" placeholder="Контактное лицо" required>
//                 <input id="register_school_name" name="school" style="margin-top: 10px;" autocomplete="off" readonly onfocus="this.removeAttribute('readonly')" class="form-control form-placeholder" type="text" placeholder="Название школы" required>
//                 <input id="register_phone" name="phone" style="margin-top: 10px;" autocomplete="off" readonly onfocus="this.removeAttribute('readonly')" class="form-control form-placeholder" type="text" placeholder="Телефон" required>
//                 <input id="register_email" name="email" style="margin-top: 10px;" autocomplete="off" readonly onfocus="this.removeAttribute('readonly')" class="form-control form-placeholder" type="email" placeholder="E-Mail" required>
//             <div>`);
//         }else if($btn.hasClass('js-reg-repetitor')){
//             $modal.find('.col-md-12').after(`
//             <div class="js-repetitor-form">
//                 <input type="hidden" name="type" value="repetitor">
//                 <input id="register_username" name="name" autocomplete="off" readonly onfocus="this.removeAttribute('readonly')" class="form-control form-placeholder" type="text" placeholder="Ваше имя" required>
//                 <input id="register_phone" name="phone" style="margin-top: 10px;" autocomplete="off" readonly onfocus="this.removeAttribute('readonly')" class="form-control form-placeholder" type="text" placeholder="Телефон" required>
//                 <input id="register_email" name="email" style="margin-top: 10px;" autocomplete="off" readonly onfocus="this.removeAttribute('readonly')" class="form-control form-placeholder" type="email" placeholder="E-Mail" required>
//             <div>`);
//         }else{
//             $modal.find('.col-md-12').after(`
//             <div class="js-user-form">
//                 <input id="register_username" name="username" autocomplete="off" readonly onfocus="this.removeAttribute('readonly')" class="form-control form-placeholder" type="text" placeholder="Ваше имя" required>
//                 <input id="register_email" name="email" autocomplete="off" readonly onfocus="this.removeAttribute('readonly')" class="form-control form-placeholder" type="text" placeholder="E-Mail" required>
//                 <input id="register_password" name="password" autocomplete="off" readonly onfocus="this.removeAttribute('readonly')" class="form-control form-placeholder" type="password" placeholder="Пароль" required>
//                 <div class="col-md-12 text-center mt-3">
//                     <a href="/auth/vk" class="no-history" onclick="ym(10436950,'reachGoal','socialreg');ga('send', 'event', 'socialreg', 'sendform');return true;"> <button type="button" class="submit-icon"><i class="fa fa-vk"></i></button></a>
//                     <a href="/auth/facebook" class="no-history" onclick="ym(10436950,'reachGoal','socialreg');ga('send', 'event', 'socialreg', 'sendform');return true;"> <button type="button" class="submit-icon"><i class="fa fa-facebook"></i></button></a>
//                     <a href="/auth/google" class="no-history" onclick="ym(10436950,'reachGoal','socialreg');ga('send', 'event', 'socialreg', 'sendform');return true;"> <button type="button" class="submit-icon"><i class="fa fa-google"></i></button></a>
//                     <a href="/auth/yandex" class="no-history" onclick="ym(10436950,'reachGoal','socialreg');ga('send', 'event', 'socialreg', 'sendform');return true;"> <button type="button" class="submit-icon"><i class="fa font-weight-bold">Я</i></button></a>
//                 </div>
//             <div>`);
//         }
//     });
// });