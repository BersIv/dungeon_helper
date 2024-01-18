package com.example.dungeon_helper.fragments.main

import android.annotation.SuppressLint
import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import com.example.dungeon_helper.MainActivity
import android.widget.TextView
import com.example.dungeon_helper.AuthActivity
import com.example.dungeon_helper.R
import com.example.dungeon_helper.databinding.FragmentAccountMainBinding
import android.content.Intent
import androidx.lifecycle.Observer
import androidx.test.internal.runner.junit4.statement.UiThreadStatement.runOnUiThread
import com.example.dungeon_helper.SharedViewModel
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
import java.io.IOException

class AccountMain : Fragment() {

    companion object {
        fun newInstance() = AccountMain()
    }
    private lateinit var sharedViewModel: SharedViewModel


    private var _binding: FragmentAccountMainBinding? = null
    private  val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View {
        val accountMainViewModel = ViewModelProvider(this)[AccountMainViewModel::class.java]

        _binding = FragmentAccountMainBinding.inflate(inflater, container, false)
        val root: View = binding.root

        val textView: TextView = binding.textAccount
        accountMainViewModel.text.observe(viewLifecycleOwner) {
            textView.text = it
        }

        return root
    }

    override fun onDestroyView() {
        super.onDestroyView()
        _binding = null
    }

    @SuppressLint("RestrictedApi")
    override fun onStart() {
        super.onStart()

        binding.nickFill.text

        sharedViewModel = ViewModelProvider(requireActivity())[SharedViewModel::class.java]

        sharedViewModel.nickname.observe(viewLifecycleOwner, Observer {
            // updating data in displayMsg
            println(it)
            binding.nickFill.text = it
        })

        sharedViewModel.email.observe(viewLifecycleOwner, Observer {
            // updating data in displayMsg
            binding.emailFill.text = it
        })


        val changePwdBtn = binding.changePwdBtn
        val exAccBtn = binding.exAccBtn
        val editBtn = binding.editBtn


        changePwdBtn.setOnClickListener {
          (activity as MainActivity).navController.navigate(R.id.action_navigation_account_to_accountRestorePwd2)
         }
        editBtn.setOnClickListener{
            (activity as MainActivity).navController.navigate(R.id.action_navigation_account_to_accountEdit)
        }
        exAccBtn.setOnClickListener {
            (requireActivity() as MainActivity).showConfirmationDialog(
                "Подтверждение выхода",
                "Вы уверены,что хотите выйти из аккаунта?",
                {
                    GlobalScope.launch(Dispatchers.Main) {

                        val client = OkHttpClient()

                        val request = Request.Builder()
                            .url("http://194.247.187.44:5000/auth/logout")
                            .post("".toRequestBody())
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
                            println(response.body!!.string())

                            val intent = Intent(activity as MainActivity, AuthActivity::class.java)
                            startActivity(intent)

                        } catch (e: IOException) {
                            println("Ошибка подключения: $e");
                        }


                    }
                },
                {
                }
            )


        }
    }
}