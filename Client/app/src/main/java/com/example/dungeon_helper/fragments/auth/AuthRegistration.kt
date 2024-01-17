package com.example.dungeon_helper.fragments.auth

import android.os.Build
import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import androidx.annotation.RequiresApi
import androidx.fragment.app.Fragment
import androidx.lifecycle.ViewModelProvider
import com.example.dungeon_helper.AuthActivity
import com.example.dungeon_helper.R
import com.example.dungeon_helper.databinding.FragmentAuthRegistrationBinding
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.GlobalScope
import kotlinx.coroutines.launch
import kotlinx.coroutines.withContext
import okhttp3.MediaType.Companion.toMediaType
import okhttp3.OkHttpClient
import okhttp3.Request
import okhttp3.RequestBody
import okhttp3.RequestBody.Companion.toRequestBody
import org.json.JSONObject
import java.io.BufferedReader
import java.io.IOException
import java.io.InputStream
import java.io.InputStreamReader
import java.io.Reader
import java.io.StringWriter
import java.io.Writer
import kotlin.io.encoding.ExperimentalEncodingApi


class AuthRegistration : Fragment() {

    companion object {
        fun newInstance() = AuthRegistration()
    }

    private lateinit var viewModel: AuthRegistrationViewModel

    private var _binding: FragmentAuthRegistrationBinding? = null
    private  val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        val authRegistrationViewModel = ViewModelProvider(this)[AuthRegistrationViewModel::class.java]

        _binding = FragmentAuthRegistrationBinding.inflate(inflater, container, false)
        val root: View = binding.root

        val textView: TextView = binding.textRegistration
        authRegistrationViewModel.text.observe(viewLifecycleOwner) {
            textView.text = it
        }

        return root
    }

    override fun onDestroyView() {
        super.onDestroyView()
        _binding = null
    }

    @OptIn(ExperimentalEncodingApi::class)
    @RequiresApi(Build.VERSION_CODES.O)
    override fun onStart() {
        super.onStart()

        val nick = binding.textFieldNick.editText
        val mail = binding.textFieldMail.editText
        val pwd = binding.textFieldPwd.editText
        val pwdRep = binding.textFieldRepPwd.editText

        val backBtn = binding.backBtn
        val regBtn = binding.registrationBtn




        backBtn.setOnClickListener {
            (activity as AuthActivity).navController.navigate(R.id.action_authRegistration_to_auth)
        }

        regBtn.setOnClickListener {
            GlobalScope.launch(Dispatchers.Main) {

                val client = OkHttpClient()
                    val jsonBody = JSONObject().apply {
                        put("email", mail?.text.toString())
                        put("password", pwd?.text.toString())
                        put("nickname", nick?.text.toString())
                        put("avatar", "dsdsd4")
                    }

                val mediaType = "application/json; charset=utf-8".toMediaType()
                val body: RequestBody = jsonBody.toString().toRequestBody(mediaType)

                val request = Request.Builder()
                    .url("http://194.247.187.44:5000/auth/registration")
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
                    println("Server: ${response.header("Server")}")
                    // вывод тела ответа
                    println(response.body!!.string())

                } catch (e: IOException) {
                    println("Ошибка подключения: $e");
                }

            //(activity as AuthActivity).navController.navigate(R.id.action_authRegistration_to_auth)
            }
        }
    }



}