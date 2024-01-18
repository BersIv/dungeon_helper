package com.example.dungeon_helper


import android.os.Bundle
import androidx.appcompat.app.AppCompatActivity
import androidx.lifecycle.ViewModelProvider
import androidx.navigation.NavController
import androidx.navigation.findNavController
import androidx.navigation.ui.AppBarConfiguration
import androidx.navigation.ui.setupActionBarWithNavController
import androidx.navigation.ui.setupWithNavController
import com.example.dungeon_helper.databinding.ActivityMainBinding
import com.example.dungeon_helper.fragments.main.AccountMain
import com.google.android.material.bottomnavigation.BottomNavigationView
import com.google.android.material.dialog.MaterialAlertDialogBuilder

/*
 *      MainActivity
 *      - opens our fragment which has the UI
 */
class MainActivity : AppCompatActivity() {

    private lateinit var binding: ActivityMainBinding

    lateinit var navController: NavController

    //private lateinit var bottomNavigationView: BottomNavigationView
    fun getNavView(): BottomNavigationView {
        return binding.navView
    }

    fun showConfirmationDialog(
        title: String,
        message: String,
        onPositiveButtonClick: () -> Unit,
        onNegativeButtonClick: () -> Unit
    ) {
        val builder = MaterialAlertDialogBuilder(this)
        builder.setTitle(title)
            .setMessage(message)
            .setPositiveButton("Да") { _, _ -> onPositiveButtonClick.invoke() }
            .setNegativeButton("Нет") { dialog, _ ->
                onNegativeButtonClick.invoke()
                dialog.dismiss()
            }
            .setCancelable(false)

        val dialog = builder.create()
        dialog.show()
    }
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)

        binding = ActivityMainBinding.inflate(layoutInflater)
        setContentView(binding.root)

        val navView: BottomNavigationView = binding.navView

        navController = findNavController(R.id.nav_host_activity_main)

        val appBarConfiguration = AppBarConfiguration(
            setOf(
                R.id.navigation_info,
                R.id.navigation_character,
                R.id.navigation_lobby,
                R.id.navigation_account
            )
        )
        setupActionBarWithNavController(navController, appBarConfiguration)
        navView.setupWithNavController(navController)

        supportActionBar?.hide()

        val receivedData = intent.getStringExtra("key")

        val sharedViewModel = ViewModelProvider(this)[SharedViewModel::class.java]

        sharedViewModel.token.value = intent.getStringExtra("token")
        sharedViewModel.email.value = intent.getStringExtra("mail")
        sharedViewModel.nickname.value = intent.getStringExtra("nick")
        sharedViewModel.avatar.value = intent.getStringExtra("avatar")

//        // Создание фрагмента и передача данных через метод
//        val fragment = AccountMain()
//        fragment.receiveData(receivedData)
//
//        // Заменить или добавить фрагмент в контейнер
//        supportFragmentManager.beginTransaction()
//            .replace(R.id.containerMain, fragment)
//            .commit()
    }
}
