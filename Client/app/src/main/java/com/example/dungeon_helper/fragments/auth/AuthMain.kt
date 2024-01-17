package com.example.dungeon_helper.fragments.auth

import android.content.Context
import android.content.Intent
import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.webkit.CookieManager
import android.widget.TextView
import androidx.fragment.app.Fragment
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.Observer
import androidx.lifecycle.ViewModelProvider
import com.example.dungeon_helper.AuthActivity
import com.example.dungeon_helper.MainActivity
import com.example.dungeon_helper.R
import com.example.dungeon_helper.SharedViewModel
import com.example.dungeon_helper.databinding.FragmentAuthMainBinding
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.GlobalScope
import kotlinx.coroutines.launch
import kotlinx.coroutines.withContext
import okhttp3.Cookie
import okhttp3.CookieJar
import okhttp3.HttpUrl
import okhttp3.Interceptor
import okhttp3.MediaType.Companion.toMediaType
import okhttp3.OkHttpClient
import okhttp3.Request
import okhttp3.RequestBody
import okhttp3.RequestBody.Companion.toRequestBody
import org.json.JSONObject
import java.io.IOException


class AuthMain : Fragment() {

    companion object {
        fun newInstance() = AuthMain()
    }

    private lateinit var sharedViewModel: SharedViewModel

    private var _binding: FragmentAuthMainBinding? = null
    private  val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        val authMainViewModel = ViewModelProvider(this)[AuthMainViewModel::class.java]

        _binding = FragmentAuthMainBinding.inflate(inflater, container, false)
        val root: View = binding.root

        val textView: TextView = binding.textAuthMain
        authMainViewModel.text.observe(viewLifecycleOwner) {
            textView.text = it
        }

        return root

    }

    override fun onDestroyView() {
        super.onDestroyView()
        _binding = null
    }

    override fun onStart() {
        super.onStart()

        sharedViewModel = ViewModelProvider(requireActivity())[SharedViewModel::class.java]

        val mail = binding.textFieldMail.editText
        val pwd = binding.textFieldPwd.editText

        val regBtn = binding.registrationBtn
        val restoreBtn = binding.restorePwdBtn
        val loginBtn = binding.loginBtn
        val loginGoogleBtn = binding.loginGoogleBtn

        regBtn.setOnClickListener {
            (activity as AuthActivity).navController.navigate(R.id.action_auth_to_authRegistration)
        }

        restoreBtn.setOnClickListener {
            (activity as AuthActivity).navController.navigate(R.id.action_auth_to_authRestorePwd)
        }

        loginBtn.setOnClickListener {
            GlobalScope.launch(Dispatchers.Main) {

                val client = OkHttpClient().newBuilder()
                    .addInterceptor(Interceptor { chain: Interceptor.Chain ->
                        val original = chain.request()
                        val authorized = original.newBuilder()
                            .addHeader("Cookie", "cookie-name=cookie-value")
                            .build()
                        chain.proceed(authorized)
                    })
                    .build()

                val jsonBody = JSONObject().apply {
                    put("email", mail?.text.toString())
                    put("password", pwd?.text.toString())
                }

                val mediaType = "application/json; charset=utf-8".toMediaType()
                val body: RequestBody = jsonBody.toString().toRequestBody(mediaType)

                val request = Request.Builder()
                    .url("http://194.247.187.44:5000/auth/byEmail")
                    .post(body)
                    .build()

                try {
                    val response = withContext(Dispatchers.IO) {
                        client.newCall(request).execute()
                    }

                    if (!response.isSuccessful) {
                        throw IOException("Запрос к серверу не был успешен:" +
                                " ${response.code} ${response.message}")
                    }
                    // пример получения конкретного заголовка ответа
                    println("${response.code} ${response.message}")
                    // вывод тела ответа
                    //println(response.body!!.string())
                    sharedViewModel.getData(response.body!!.string())

                    val cookies = response.headers("Set-Cookie")
                    if (cookies.isNotEmpty()) {
                        val sharedPreferences = requireContext().getSharedPreferences("Cookies", Context.MODE_PRIVATE)
                        val editor = sharedPreferences.edit()
                        editor.putString("Cookie", cookies.joinToString(";"))
                        editor.apply()

                        sharedViewModel.token.value = getCookieValue()

                        val intent = Intent(activity as AuthActivity, MainActivity::class.java)
                        intent.putExtra("token", sharedViewModel.getToken().toString())
                        intent.putExtra("mail", sharedViewModel.getEmail().toString())
                        intent.putExtra("nick", sharedViewModel.getNick().toString())
                        intent.putExtra("avatar", sharedViewModel.getAvatar().toString())
                        startActivity(intent)
                    }



                } catch (e: IOException) {
                    println("Ошибка подключения: $e");
                }


            }
        }

        loginGoogleBtn.setOnClickListener {

        }

    }

    fun getCookieValue(): String {
        val sharedPreferences = requireContext().getSharedPreferences("Cookies", Context.MODE_PRIVATE)
        val savedCookie = sharedPreferences.getString("Cookie", "").toString()

        val cookiePairs = savedCookie.split(";")
        if (cookiePairs.isNotEmpty()) {
            for (cookiePair in cookiePairs) {
                val pair = cookiePair.trim().split("=")
                if (pair.size == 2) {
                    // Мы убеждаемся, что у нас есть оба элемента перед получением значения
                    return "Bearer " + pair[1]
                }
            }
        }

        // Возвращаем пустую строку, если не удалось получить значение
        return ""
    }

}